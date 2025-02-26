from pydantic import BaseModel, Field, PydanticUserError
import os
import json
from dotenv import load_dotenv
import logging
from typing import List, Optional, Dict, Any

load_dotenv('.env')

def get_standard_options(option_type: str) -> list[dict]:
    """Get standardized option items for common option types."""
    options = {
        "Meat Temperature": [
            {"name": "Well Done", "price": 0},
            {"name": "Medium-Well", "price": 0},
            {"name": "Medium", "price": 0},
            {"name": "Medium-Rare", "price": 0},
            {"name": "Rare", "price": 0}
        ],
        "Choice of Side": [
            {"name": "French Fries", "price": 0},
            {"name": "Sweet Potato Fries", "price": 0},
            {"name": "Onion Rings", "price": 0},
            {"name": "Fruit", "price": 0},
            {"name": "Homemade Chips", "price": 0},
            {"name": "Pasta Salad", "price": 0},
            {"name": "Potato Salad", "price": 0}
        ],
        "Choice of Bread": [
            {"name": "White", "price": 0},
            {"name": "Wheat", "price": 0},
            {"name": "Flour Wrap", "price": 0},
            {"name": "Garlic Wrap", "price": 0},
            {"name": "Gluten Free Wrap", "price": 1}
        ],
        "Choice of Cheese": [
            {"name": "American", "price": 0},
            {"name": "Cheddar", "price": 0},
            {"name": "Swiss", "price": 0},
            {"name": "Provolone", "price": 0},
            {"name": "Pepperjack", "price": 0},
            {"name": "No Cheese", "price": 0}
        ],
        "Wings Sauce": [
            {"name": "Buffalo", "price": 0},
            {"name": "BBQ", "price": 0},
            {"name": "Blazing BBQ", "price": 0},
            {"name": "Incinerator", "price": 0},
            {"name": "Garlic Parm", "price": 0}
        ],
        "Wings Dipping Sauce": [
            {"name": "Ranch", "price": 0},
            {"name": "Blue Cheese", "price": 0}
        ],
        "Salad Dressing": [
            {"name": "Balsamic", "price": 0},
            {"name": "Caesar", "price": 0},
            {"name": "Ranch", "price": 0},
            {"name": "Blue Cheese", "price": 0},
            {"name": "Honey Mustard", "price": 0},
            {"name": "Italian", "price": 0},
            {"name": "Thousand Island", "price": 0}
        ]
    }
    return options.get(option_type, [])

def has_item_options(item_name: str, item_options: list, item_description: str = "") -> bool:
    """Determine if an item needs options based on its name, existing options, and description."""
    # Terminal items never have options
    item_name_lower = item_name.lower()
    
    # Check known terminal items first
    if any([
        "side" in item_name_lower,
        "calamari" in item_name_lower,
        "mozzarella sticks" in item_name_lower,
        "par three platter" in item_name_lower,
        "fish and chips" in item_name_lower,
        "bacon mac" in item_name_lower,
        "loaded skins" in item_name_lower,
        "stuffed mushrooms" in item_name_lower,
        "soup of the day" in item_name_lower,
        "side house" in item_name_lower,
        "side caesar" in item_name_lower,
        "bang bang tempura" in item_name_lower,
        "loaded quesadilla" in item_name_lower,
        "bbq pork quesadilla" in item_name_lower
    ]):
        return False
    
    # Check for existing options
    if len(item_options) > 0:
        return True
    
    # Check for items that should have options based on description
    if "choice of side" in item_description.lower():
        # Special cases that don't need options despite description
        if any([
            "fish and chips" in item_name_lower,
            "bang bang tempura" in item_name_lower
        ]):
            return False
        return True
    
    # Check item type
    return any([
        ("burger" in item_name_lower and "portobello" not in item_name_lower),
        ("steak" in item_name_lower and "tuna" not in item_name_lower),
        ("sandwich" in item_name_lower and "side" not in item_name_lower),
        "deli" in item_name_lower,
        "wings" in item_name_lower,
        ("salad" in item_name_lower and "side" not in item_name_lower 
         and "potato" not in item_name_lower and "pasta" not in item_name_lower),
        "fajitas" in item_name_lower
    ])

def get_option_min_max(option_name: str, item_name: str = "") -> tuple[int, int]:
    """Get min/max values for option types."""
    # Remove Options are always optional
    if option_name == "Remove Options":
        return (0, 0)
    
    # JP's Ace has optional sides
    if option_name == "Choice of Side" and "JP's Ace" in item_name:
        return (0, 0)
    
    # Required choices
    if option_name in ["Choice of Side", "Choice of Meat", "Choice of Bread", 
                      "Choice of Cheese", "Meat Temperature", "Wings Sauce",
                      "Wings Dipping Sauce", "Salad Dressing"]:
        return (1, 1)
    
    # Default for other options (like Extras)
    return (0, 1)

def create_go_seed_file(menu_data, club_name: str, club_address: str):
    """Create a Go seed file from the menu data."""
    go_code = []
    
    # Import statements
    go_code.append('package main\n')
    go_code.append('import (')
    go_code.append('\t"log"')
    go_code.append('\t"swoop/locations"')
    go_code.append('\t"swoop/pkg/config"')
    go_code.append('\tdatabase "swoop/pkg/db"')
    go_code.append(')\n')
    
    # Add config initialization
    go_code.append('func main() {')
    go_code.append('\terr := config.Init()')
    go_code.append('\tif err != nil {')
    go_code.append('\t\tlog.Panicln(err)')
    go_code.append('\t}')
    go_code.append('\tlog.Println("Seeding Database")')
    go_code.append('\tdb := database.Connect(config.DB())')
    go_code.append('\tapi := locations.NewAPI(db)\n')
    
    # Main function
    go_code.append('func main() {')
    go_code.append('\t// Create menu items')
    go_code.append('\tlog.Println("Seeding Database")\n')
    
    # Process locations
    for location in menu_data.locations:
        location_var = sanitize_location_var(location.name)
        go_code.append(f'\t// Seed location')
        go_code.append(f'\tlog.Println("Seeding *Location* data")')
        go_code.append(f'\t{location_var} := api.CreateLocation("{location.name}", "{location.address}")\n')
        
        # Process menus
        for menu in location.menus:
            menu_var = sanitize_name_var(menu.name)
            go_code.append(f'\t// Seed menu')
            go_code.append(f'\tlog.Println("Seeding *Menu* data")')
            go_code.append(f'\t{location_var}{menu_var} := api.CreateMenu("{menu.name}", "{menu.name}", {location_var}.ID)\n')
            
            # Process categories
            for category in menu.categories:
                category_var = sanitize_name_var(category.name)
                full_category_var = f"{location_var}{menu_var}{category_var}"
                go_code.append(f'\t// Seed category')
                go_code.append(f'\tlog.Println("Seeding *Category* data")')
                go_code.append(f'\t{full_category_var} := api.CreateCategory("{category.name}", "{category.name}", {location_var}{menu_var}.ID)\n')
                
                # Process items
                for item in category.items:
                    # Check if item has options
                    has_options = has_item_options(item.name, item.options, item.description)
                    needs_variable = False
                    
                    # Determine if item needs a variable assignment
                    if has_options:
                        # Check for items that need variables despite no current options
                        if "choice of side" in item.description.lower():
                            if not any([
                                "fish and chips" in item.name.lower(),
                                "bang bang tempura" in item.name.lower()
                            ]):
                                needs_variable = True
                        # Items with actual options need variables
                        if len(item.options) > 0:
                            needs_variable = True
                    
                    # Generate item creation code
                    go_code.append(f'\t// Seed item')
                    if needs_variable:
                        # Create variable for items that need options
                        item_var = sanitize_name_var(item.name, f"{location_var}{menu_var}{category_var}")
                        go_code.append(f'\t{item_var} := api.CreateItem("{item.name}", "{item.description}", {item.price}, {full_category_var}.ID)\n')
                    else:
                        # Use _ assignment for terminal items
                        go_code.append(f'\t_ = api.CreateItem("{item.name}", "{item.description}", {item.price}, {full_category_var}.ID)\n')
                    
                    # Process options in standard order
                    option_order = ["Meat Temperature", "Choice of Side", "Choice of Meat", 
                                  "Choice of Bread", "Choice of Cheese", "Wings Sauce",
                                  "Wings Dipping Sauce", "Salad Dressing", "Remove Options"]
                    
                    # Add standard options based on item type
                    if "burger" in item.name.lower() or "steak" in item.name.lower():
                        standard_options = ["Meat Temperature", "Choice of Side"]
                    elif "sandwich" in item.name.lower() or "deli" in item.name.lower():
                        standard_options = ["Choice of Side", "Choice of Bread"]
                    elif "wings" in item.name.lower():
                        standard_options = ["Wings Sauce", "Wings Dipping Sauce"]
                    elif "salad" in item.name.lower() and not "side" in item.name.lower():
                        standard_options = ["Salad Dressing"]
                    else:
                        standard_options = []
                    
                    # Add standard options first
                    for opt_type in option_order:
                        if opt_type in standard_options:
                            min_val, max_val = get_option_min_max(opt_type, item.name)
                            opt_items = get_standard_options(opt_type)
                            if opt_items:
                                go_code.append(f'\t// Add {opt_type}')
                                go_code.append(f'\t_ = api.CreateOption("{opt_type}", "{opt_type}", {min_val}, {max_val}, {item_var}.ID, []locations.OptionItem{{')
                                for opt_item in opt_items:
                                    go_code.append(f'\t\tlocations.OptionItem{{Name: "{opt_item["name"]}", Description: "{opt_item["name"]}", Price: {opt_item["price"]}}},')
                                go_code.append('\t})\n')
                    
                    # Process custom options
                    for option in item.options:
                        if option.name not in standard_options:
                            min_val, max_val = get_option_min_max(option.name, item.name)
                            go_code.append(f'\t// Add option')
                            go_code.append(f'\t_ = api.CreateOption("{option.name}", "{option.name}", {min_val}, {max_val}, {item_var}.ID, []locations.OptionItem{{')
                            for opt_item in option.option_items:
                                go_code.append(f'\t\tlocations.OptionItem{{Name: "{opt_item.name}", Description: "{opt_item.name}", Price: {opt_item.price}}},')
                            go_code.append('\t})\n')
    
    go_code.append('}')
    
    # Write to file
    output_directory = "seed_files"
    os.makedirs(output_directory, exist_ok=True)
    go_file_path = os.path.join(output_directory, "gatesFour.go")
    
    with open(go_file_path, "w", encoding="utf-8") as f:
        f.write("\n".join(go_code))
    
    print(f"Generated Go seed file at: {go_file_path}")

def sanitize_location_var(name: str) -> str:
    """Convert location name to a valid Go variable name."""
    return "".join(x for x in name.title().replace(" ", "") if x.isalnum())

def sanitize_name_var(name: str, prefix: str = "") -> str:
    """Convert menu/category/item name to a valid Go variable name.
    Args:
        name: The name to sanitize
        prefix: Optional prefix to ensure uniqueness across menu levels
    """
    sanitized = "".join(x for x in name.title().replace(" ", "") if x.isalnum())
    return f"{prefix}{sanitized}" if prefix else sanitized
