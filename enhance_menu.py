import os
import json
import google.generativeai as genai
from dotenv import load_dotenv
import logging
import time
from typing import Union, List, Dict, Optional

# Setup logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# Load API key from .env file
load_dotenv('.env')
api_key = os.getenv("GOOGLE_GEMINI_API")

# Configure the model
genai.configure(api_key=api_key)
model = genai.GenerativeModel("gemini-2.0-flash")

def extract_ingredients(description: str) -> list[str]:
    """Extract ingredients from item description for generating remove options."""
    common_ingredients = [
        "lettuce", "tomato", "onion", "cheese", "bacon", "ham", "turkey", "chicken",
        "beef", "mushroom", "pepper", "jalapeno", "mayo", "mustard", "sauce"
    ]
    return [ingredient for ingredient in common_ingredients if ingredient in description.lower()]

def generate_remove_options(description: str, required_choices: Optional[List[str]] = None) -> List[Dict]:
    """Generate Remove Options from item description, excluding required choices."""
    if required_choices is None:
        required_choices = []
    
    ingredients = extract_ingredients(description)
    remove_options = []
    
    # Group ingredients by type for ordering
    ingredient_groups = {
        'proteins': ['bacon', 'ham', 'turkey', 'chicken', 'beef'],
        'dairy': ['cheese', 'american', 'cheddar', 'swiss', 'provolone', 'pepperjack'],
        'vegetables': ['lettuce', 'tomato', 'onion', 'mushroom', 'pepper', 'jalapeno'],
        'condiments': ['mayo', 'mustard', 'sauce']
    }
    
    # Sort ingredients by group
    for group in ['proteins', 'dairy', 'vegetables', 'condiments']:
        for ingredient in ingredient_groups[group]:
            if ingredient in ingredients and ingredient not in required_choices:
                name = ingredient.capitalize()
                remove_options.append({
                    "name": f"No {name}",
                    "description": f"No {name}",
                    "price": 0
                })
    
    return remove_options

def enhance_menu_item(item: dict) -> dict:
    """Enhance a single menu item with proper options."""
    enhanced_item = item.copy()
    enhanced_item["options"] = []
    
    # Define standard option order
    option_order = [
        "Meat Temperature",  # Must come first for burgers
        "Choice of Side",    # Second for items with sides
        "Choice of Meat",    # Third for deli items
        "Choice of Bread",   # Fourth for deli items
        "Choice of Cheese",  # Fifth for deli items
        "Wings Sauce",       # For wings
        "Wings Dipping Sauce", # For wings
        "Salad Dressing",    # For salads
        "Extras",           # For salads
        "Remove Options"     # Always last
    ]
    
    # Determine required options based on item type and description
    required_options = []
    
    # Define items that need specific options
    item_name = item["name"].lower().strip()
    
    # Items that need Choice of Side
    side_items = {
        "ahi tuna steak", "bang bang tempura", "chicken fajitas", "✓ portobello burger",
        "irongate deli", "philly cheesesteak", "the italian", "*blackened salmon sandwich",
        "jp's club sandwich", "buffalo chicken sandwich", "*gates four burger",
        "*big sky burger", "*carolina burger", "jp's ace"
    }
    
    # Items that need Meat Temperature
    temp_items = {
        "*gates four burger", "*big sky burger", "*carolina burger"
    }
    
    # Items that need specific options
    if item_name in side_items:
        required_options.append("Choice of Side")
    
    if item_name in temp_items or "✓ portobello burger" in item_name:
        required_options.append("Meat Temperature")
    
    # Check for specific items
    if "irongate deli" in item_name:
        required_options.extend(["Choice of Meat", "Choice of Bread", "Choice of Cheese"])
    
    if "chicken wings" in item_name:
        required_options.extend(["Wings Sauce", "Wings Dipping Sauce"])
    
    if "salad" in item_name:
        required_options.append("Salad Dressing")
        required_options.append("Salad Dressing")
    
    # Add options in standard order
    for option_type in option_order:
        if option_type in required_options:
            # Get min/max values based on option type and item
            min_val = 0 if option_type == "Remove Options" or (option_type == "Choice of Side" and "jp's ace" in item["name"].lower()) else 1
            max_val = 0 if option_type == "Remove Options" or (option_type == "Choice of Side" and "jp's ace" in item["name"].lower()) else 1
            
            option_items = get_standard_options(option_type, item["name"])
            if option_items:
                enhanced_item["options"].append({
                    "name": option_type,
                    "min": min_val,
                    "max": max_val,
                    "optionItems": option_items
                })
    
    # Add Remove Options last
    remove_options = generate_remove_options(item["description"])
    if remove_options:
        enhanced_item["options"].append({
            "name": "Remove Options",
            "min": 0,
            "max": 0,
            "optionItems": remove_options
        })
    
    return enhanced_item

def get_standard_options(option_type: str, item_name: str = "") -> list[dict]:
    """Get standardized option items for common option types."""
    options = {
        "Wings Sauce": [
            {"name": "Buffalo", "description": "Buffalo", "price": 0},
            {"name": "BBQ", "description": "BBQ", "price": 0},
            {"name": "Blazing BBQ", "description": "Blazing BBQ", "price": 0},
            {"name": "Incinerator", "description": "Incinerator", "price": 0},
            {"name": "Garlic Parm", "description": "Garlic Parm", "price": 0}
        ],
        "Wings Dipping Sauce": [
            {"name": "Ranch", "description": "Ranch", "price": 0},
            {"name": "Blue Cheese", "description": "Blue Cheese", "price": 0}
        ],
        "Salad Dressing": [
            {"name": "Balsalmic", "description": "Balsalmic", "price": 0},
            {"name": "Caesar", "description": "Caesar", "price": 0},
            {"name": "Ranch", "description": "Ranch", "price": 0},
            {"name": "Bleu Cheese", "description": "Bleu Cheese", "price": 0},
            {"name": "Honey Mustard", "description": "Honey Mustard", "price": 0},
            {"name": "Italian", "description": "Italian", "price": 0},
            {"name": "Thousand Island", "description": "Thousand Island", "price": 0},
            {"name": "Balsalmic Vinaigrette", "description": "Balsalmic Vinaigrette", "price": 0}
        ],
        "Choice of Meat": [
            {"name": "Ham", "description": "Ham", "price": 0},
            {"name": "Turkey", "description": "Turkey", "price": 0},
            {"name": "Roast Beef", "description": "Roast Beef", "price": 0},
            {"name": "Capicola", "description": "Capicola", "price": 0}
        ],
        "Choice of Bread": [
            {"name": "White", "description": "White", "price": 0},
            {"name": "Wheat", "description": "Wheat", "price": 0},
            {"name": "Flour Wrap", "description": "Flour Wrap", "price": 0},
            {"name": "Garlic Wrap", "description": "Garlic Wrap", "price": 0},
            {"name": "Gluten Free Wrap", "description": "Gluten Free Wrap", "price": 1}
        ],
        "Choice of Cheese": [
            {"name": "American", "description": "American", "price": 0},
            {"name": "Cheddar", "description": "Cheddar", "price": 0},
            {"name": "Swiss", "description": "Swiss", "price": 0},
            {"name": "Provolone", "description": "Provolone", "price": 0},
            {"name": "Pepperjack", "description": "Pepperjack", "price": 0},
            {"name": "No Cheese", "description": "No Cheese", "price": 0}
        ],
        "Choice of Side": [
            {"name": "French Fries", "description": "French Fries", "price": 2 if "jp's ace" in item_name.lower() else 0},
            {"name": "Sweet Potato Fries", "description": "Sweet Potato Fries", "price": 2 if "jp's ace" in item_name.lower() else 0},
            {"name": "Onion Rings", "description": "Onion Rings", "price": 2 if "jp's ace" in item_name.lower() else 0},
            {"name": "Fruit", "description": "Fruit", "price": 2 if "jp's ace" in item_name.lower() else 0},
            {"name": "Homemade Chips", "description": "Homemade Chips", "price": 2 if "jp's ace" in item_name.lower() else 0},
            {"name": "Pasta Salad", "description": "Pasta Salad", "price": 2 if "jp's ace" in item_name.lower() else 0},
            {"name": "Potato Salad", "description": "Potato Salad", "price": 2 if "jp's ace" in item_name.lower() else 0}
        ],
        "Meat Temperature": [
            {"name": "Well Done", "description": "Well Done", "price": 0},
            {"name": "Medium-Well", "description": "Medium-Well", "price": 0},
            {"name": "Medium", "description": "Medium", "price": 0},
            {"name": "Medium-Rare", "description": "Medium-Rare", "price": 0},
            {"name": "Rare", "description": "Rare", "price": 0}
        ]
    }
    return options.get(option_type, [])

def enhance_menu():
    """Enhance the menu with proper options and structure."""
    # Load the original menu
    club_name = os.getenv("CLUB_NAME", "gatesFour").replace(" ", "_").lower()
    input_file_path = os.path.join("output_files", f"extracted_menu_{club_name}.json")
    
    with open(input_file_path, "r", encoding="utf-8") as infile:
        menu_data = json.load(infile)
    
    # Create enhanced menu structure
    enhanced_menu = {
        "locations": [{
            "name": "Gates Four Golf & Country Club",
            "address": "Fayetteville, NC",
            "menus": [{
                "name": "JPs Bar and Grill",
                "categories": []
            }]
        }]
    }
    
    # Process each category
    for category in menu_data["menu"]["categories"]:
        enhanced_category = {
            "name": category["name"],
            "items": []
        }
        
        # Process each item
        for item in category["items"]:
            enhanced_item = enhance_menu_item(item)
            enhanced_category["items"].append(enhanced_item)
        
        enhanced_menu["locations"][0]["menus"][0]["categories"].append(enhanced_category)
    
    # Save enhanced menu
    output_file_path = os.path.join("output_files", f"enhanced_menu_{club_name}.json")
    with open(output_file_path, "w", encoding="utf-8") as outfile:
        json.dump(enhanced_menu, outfile, indent=2)
    
    print(f"Enhanced Menu saved to: {output_file_path}")

if __name__ == "__main__":
    enhance_menu()
