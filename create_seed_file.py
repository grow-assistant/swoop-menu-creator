from google import genai
from pydantic import BaseModel, Field, PydanticUserError
import os
from dotenv import load_dotenv
import logging
import json
from typing import List, Optional

load_dotenv('.env')
api_key = os.getenv("GOOGLE_GEMINI_API")  # Set your API key here

# Create a Gemini client
client = genai.Client(api_key=api_key)

# Define the model you are going to use
model_id = "gemini-2.0-flash"  # or use other available model names

# No need for extract_structured_data, we are loading from JSON
# --- Pydantic models for the menu. --- (Keep these, they are still useful for type hinting)
class MenuItemOption(BaseModel):
    name: str = Field(description="The name of the menu item option")
    price: float = Field(description="The price for the option")
    description: str = Field(description="The description of the option")

class Option(BaseModel):
    name: str = Field(description="Name of the option category")
    min: int = Field(description="Minimum number of options that must be selected")
    max: int = Field(description="Maximum number of options that can be selected")
    optionItems: List[MenuItemOption] = Field(description="List of option items")

class MenuItem(BaseModel):
    name: str = Field(description="Name of the menu item")
    description: str = Field(description="Description of the menu item")
    price: float = Field(description="Price of the menu item")
    options: Optional[List[Option]] = Field(default=None, description="Options for the menu item")

class Category(BaseModel):
    name: str = Field(description="Name of the category")
    items: List[MenuItem] = Field(description="List of menu items in the category")

class Menu(BaseModel):
    name: str = Field(description="Name of the menu")
    categories: List[Category] = Field(description="List of categories in the menu")

class Location(BaseModel):
    name: str = Field(description="Name of the location")
    address: str = Field(description="Address of the location")
    menus: List[Menu] = Field(description="List of menus at the location")

class MenuStructure(BaseModel):
    locations: List[Location] = Field(description="List of locations")

# --- Helper function to sanitize names for Go variables ---
def sanitize_location_var(name):
    # Changed to match gatesFour.go's camelCase naming without underscores
    name = name.replace('&', 'and').replace("'", "")
    # Remove all non-alphanumeric except spaces
    name = ''.join([c if c.isalnum() or c.isspace() else '' for c in name])
    words = name.split()
    if not words:
        return ''
    # Lowercase first word
    words[0] = words[0].lower()
    # Capitalize remaining words and remove spaces
    return words[0] + ''.join(word.capitalize() for word in words[1:])

def sanitize_name_var(name):
    # Match the shorter variable names from gatesFour.go
    name = name.replace('&', 'and').replace(" Lunch Menu", "").replace("'", "")
    # Remove special characters and split into words
    name = ''.join([c if c.isalnum() or c.isspace() else '' for c in name])
    words = name.split()
    return ''.join(word.capitalize() for word in words)

# --- Load the enhanced menu JSON ---
def load_enhanced_menu(file_path: str) -> MenuStructure:
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            data = json.load(f)
            return MenuStructure(**data)
    except (FileNotFoundError, json.JSONDecodeError, PydanticUserError) as e:
        logging.error(f"Error loading or parsing menu JSON: {e}")
        raise

# For menu generation - match the shorter menu name format
def generate_menu_section(menu):
    menu_name = menu['name'].replace(" Lunch Menu", "")  # Remove "Lunch Menu" suffix
    return f"gatesFourGolfandCountryClub{menu_name.replace(' ', '').replace('&', 'And')} := api.CreateMenu(\"{menu['name']}\", \"{menu['name']}\", gatesFourGolfandCountryClub.ID)"

# For price formatting - use integer prices instead of floats
def format_price(price):
    return f"{int(price)}"  # Remove decimal formatting

# For option items - match the style without trailing commas
def generate_option_items(items):
    option_lines = []
    for item in items:
        option_lines.append(f"locations.OptionItem{{Name: \"{item['name']}\", Description: \"{item['description']}\", Price: {format_price(item['price'])}}}")
    return '\n'.join([f"\t\t{line}" for line in option_lines])

# Fix special character handling in descriptions
def sanitize_description(desc):
    return desc.replace("'", "’")  # Use proper apostrophe

# --- Main function to generate the Go seed file ---
def create_go_seed_file(menu_data: MenuStructure, club_name: str, club_address: str):
    # --- Build the Go code string ---
    go_code = f'''package main

import (
	"log"

	"swoop/locations"
	"swoop/pkg/config"
	database "swoop/pkg/db"
)

func main() {{
	err := config.Init()
	if err != nil {{
		log.Panicln(err)
	}}

	log.Println("Seeding Database")
	db := database.Connect(config.DB())
	api := locations.NewAPI(db)

'''
    # Iterate through locations, menus, categories, items, and options
    for location in menu_data.locations:
        location_var = sanitize_location_var(location.name)
        go_code += f'''	// seed locations
	log.Println("Seeding *Location* data")
	{location_var} := api.CreateLocation("{location.name}", "{location.address}")

'''
        for menu in location.menus:
            menu_var = sanitize_name_var(menu.name)
            go_code += f'''	log.Println("Seeding *Menus* data")
	// seed menus
	{location_var}{menu_var} := api.CreateMenu("{menu.name}", "{menu.name}", {location_var}.ID)
	log.Println({location_var}{menu_var})

'''
            for category in menu.categories:
                category_var = sanitize_name_var(category.name)
                # For the category we use the location + menu + category (this ensures it matches the sample)
                full_category_var = f"{location_var}{menu_var}{category_var}"
                go_code += f'''	// seed categories
	log.Println("Seeding *Categories* data")
	{full_category_var} := api.CreateCategory("{category.name}", "{category.name}", {location_var}{menu_var}.ID)
	log.Println({full_category_var})
'''
                # Print a fixed header for items in the category
                go_code += f'''	// seed item
	log.Println("Seeding *Items* data")
'''
                # Iterate through items in the category
                for i, item in enumerate(category.items):
                    item_var = sanitize_name_var(item.name)
                    # For one-line seed items (like "Calamari")
                    if i % 3 == 0:
                        go_code += f'	_ = api.CreateItem("{item.name}", "{item.description}", {item.price}, {full_category_var}.ID)\n'
                    else:
                        # For multi-line seed items (like "Chicken Wings") we create a variable that omits the menu portion,
                        # matching the sample (e.g. Gates_Four_Golf___Country_ClubAppetizersChicken_Wings)
                        item_var_full = f"{location_var}{category_var}{item_var}"
                        go_code += f'	{item_var_full} := api.CreateItem("{item.name}", "{item.description}", {item.price}, {full_category_var}.ID)\n'
                        if item.options:
                            go_code += f'	// seed item options\n	log.Println("Seeding *Options* data")\n'
                            for option in item.options:
                                option_items_lines = ""
                                if option.optionItems:
                                    for opt_item in option.optionItems:
                                        option_items_lines += f'		locations.OptionItem{{Name: "{opt_item.name}", Description: "{opt_item.description}", Price: {opt_item.price}}},\n'
                                go_code += f'''	_ = api.CreateOption("{option.name}", "{option.name}", {option.min}, {option.max}, {item_var_full}.ID, []locations.OptionItem{{
{option_items_lines}	}})
'''
                go_code += "\n"

    go_code += '''	// seed orders
	log.Println("Seeding Completed")
}
'''

    # --- Write the Go code to the seed_files directory ---
    output_directory = "seed_files"
    os.makedirs(output_directory, exist_ok=True)
    # Use the sanitized club name for the file name, e.g. "Gates_Four_Golf___Country_Club.go"
    go_file_path = os.path.join(output_directory, f"{sanitize_location_var(club_name)}.go")
    with open(go_file_path, "w", encoding="utf-8") as f:
        f.write(go_code)

    print(f"Generated Go seed file at: {go_file_path}")

if __name__ == "__main__":
    # Replace with the actual path to your enhanced menu JSON file
    enhanced_menu_path = "output_files/improved_enhanced_menu_gates_four_golf_&_country_club.json"
    # Club name and address (typically loaded from customer_info.json)
    club_name = "Gates Four Golf & Country Club"  # Get from customer_info.json
    club_address = "Fayetteville, NC"  # Get from customer_info.json

    # Load the enhanced menu data
    menu_data = load_enhanced_menu(enhanced_menu_path)

    # Create the Go seed file
    create_go_seed_file(menu_data, club_name, club_address)
