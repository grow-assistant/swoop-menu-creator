import os
import json
import logging
from typing import Dict, List, Any

# Setup logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def load_json_file(filepath: str) -> Dict:
    """Load and parse a JSON file."""
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            return json.load(f)
    except Exception as e:
        logging.error(f"Error reading {filepath}: {e}")
        exit(1)

def compare_option_structure(enhanced_menu: Dict, reference_file: str) -> List[str]:
    """Compare option structures between enhanced menu and reference."""
    differences = []
    
    # Check for required options
    required_options = {
        "wings": ["Wings Sauce", "Wings Dipping Sauce"],
        "salad": ["Salad Dressing"],
        "burger": ["Meat Temperature", "Choice of Side"],
        "sandwich": ["Choice of Side"],
        "deli": ["Choice of Meat", "Choice of Bread", "Choice of Cheese"]
    }
    
    # Check each menu item
    for category in enhanced_menu["locations"][0]["menus"][0]["categories"]:
        for item in category["items"]:
            item_options = {opt["name"] for opt in item.get("options", [])}
            
            # Check required options based on item type
            for item_type, required_opts in required_options.items():
                if item_type in item["name"].lower():
                    missing = [opt for opt in required_opts if opt not in item_options]
                    if missing:
                        differences.append(f"Missing required options for {item['name']}: {', '.join(missing)}")
    
    return differences

def critique_menu():
    """Compare enhanced menu with reference structure."""
    # Load enhanced menu
    club_name = os.getenv("CLUB_NAME", "gatesFour").replace(" ", "_").lower()
    enhanced_menu_path = os.path.join("output_files", f"enhanced_menu_{club_name}.json")
    enhanced_menu = load_json_file(enhanced_menu_path)
    
    # Compare structures
    differences = compare_option_structure(enhanced_menu, "seed_files/gatesFour.go")
    
    if differences:
        print("\nMenu Structure Differences:")
        for diff in differences:
            print(f"- {diff}")
    else:
        print("\nMenu structure matches reference.")

if __name__ == "__main__":
    critique_menu()

