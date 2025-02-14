import os
import json
from google import genai
from dotenv import load_dotenv
import logging
import time

# Setup logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# Load API key from .env file
load_dotenv('.env')
api_key = os.getenv("GOOGLE_GEMINI_API")
client = genai.Client(api_key=api_key)

# Select your Gemini model
model_id = "gemini-2.0-flash"

# Retrieve the club name from environment variable
club_name = os.getenv("CLUB_NAME", "gatesFour").replace(" ", "_").lower()

# Load the original menu JSON from file using the club-specific filename
input_file_path = os.path.join("output_files", f"extracted_menu_{club_name}.json")
with open(input_file_path, "r", encoding="utf-8") as infile:
    original_menu = infile.read()

# Define the detailed target structure (as before, truncated for brevity)
detailed_structure = r"""
{
  "locations": [
    {
      "name": "Gates Four Golf & Country Club",
      "address": "Fayetteville, NC",
      "menus": [
        {
          "name": "JPs Bar and Grill",
          "categories": [
            {
              "name": "Appetizers",
              "items": [
                {
                  "name": "Calamari",
                  "description": "A mix of tentacles and rings, lightly coated and fried, served with a side of cocktail.",
                  "price": 12
                },
                {
                  "name": "Chicken Wings",
                  "description": "Fried jumbo chicken wings, your choice of six or twelve, tossed in one of our signature sauces: Buffalo, BBQ, Blazing BBQ, Incinerator, Garlic Parm.",
                  "price": 7,
                  "options": [
                    {
                      "name": "Wings Sauce",
                      "min": 1,
                      "max": 1,
                      "optionItems": [
                        { "name": "Buffalo", "description": "Buffalo", "price": 0 },
                        { "name": "BBQ", "description": "BBQ", "price": 0 },
                        { "name": "Blazing BBQ", "description": "Blazing BBQ", "price": 0 },
                        { "name": "Incinerator", "description": "Incinerator", "price": 0 },
                        { "name": "Garlic Parm", "description": "Garlic Parm", "price": 0 }
                      ]
                    },
                    {
                      "name": "Wings Dipping Sauce",
                      "min": 1,
                      "max": 1,
                      "optionItems": [
                        { "name": "Ranch", "description": "Ranch", "price": 0 },
                        { "name": "Blue Cheese", "description": "Blue Cheese", "price": 0 }
                      ]
                    },
                    {
                      "name": "Options",
                      "min": 1,
                      "max": 1,
                      "optionItems": [
                        { "name": "Six Wings", "description": "Six Wings", "price": 0 },
                        { "name": "Twelve Wings", "description": "Twelve Wings", "price": 5 }
                      ]
                    }
                  ]
                }
                // ... (additional items and categories)
              ]
            }
            // ... (other categories, menus, and locations)
          ]
        }
      ]
    }
  ]
}
"""

# Construct the Gemini prompt with improved instructions for an accurate output
prompt = (
    "Transform the following extracted menu JSON into a detailed menu structure that exactly matches the target JSON structure provided below.\n\n"
    "Target Menu JSON Structure (must match exactly, including all keys, nesting, and data types):\n"
    f"{detailed_structure}\n\n"
    "Extracted Menu JSON:\n"
    f"{original_menu}\n\n"
    "For each menu item, analyze the description and perform the following actions:\n"
    "1. If the item specifies a required choice (such as a choice of meat, bread, or other toppings), add a corresponding options block (e.g. 'Choice of Meat', 'Choice of Bread', or 'Add Options') that lists all valid alternatives with their option items.\n"
    "2. If the item permits ingredient removals, add a 'Remove Options' block containing common removals (such as 'No Tomatoes', 'No Onions', 'No Lettuce', 'No Mushrooms', etc.) based on the ingredients mentioned. The 'Remove Options' block should always have min=0 and max=0 to indicate these are optional removals.\n"
    "3. Ensure there is no overlap between the options provided for adding ingredients and those provided for removals. In other words, an ingredient should not appear in both the add and the remove options for the same item.\n"
    "4. If the item includes meat that typically requires a temperature specification (for example, hamburgers, steaks, or other meat-based items), include a 'Meat Temperature' or 'Cooking Preference' options block listing common cooking instructions (such as Well Done, Medium, Medium-Rare, and Rare). Ensure these options remain distinct from any removal options.\n"
    "5. Additionally, for items that require a complete set of option choices (for example, sandwiches like 'Irongate Deli' or 'JP's Club Sandwich'), include the following option blocks exactly as shown in the target structure:\n"
    "   - For sandwiches, add a 'Meat Choice' block (e.g. with options 'Ham', 'Turkey', 'Roast Beef'), a 'Bread/Wrap Choice' block (e.g. 'White Bread', 'Wheat Bread', 'Wrap' or 'Rye Bread'), and a 'Cheese Choice' block (e.g. 'Cheddar', 'Swiss', 'Provolone').\n"
    "   - For burgers, ensure that a 'Meat Temperature' block is included with options 'Rare', 'Medium Rare', 'Medium', 'Medium Well', and 'Well Done'.\n"
    "   - Ensure that no ingredient appears in both an addition options block and a removal options block.\n\n"
    "6. Make the transformation dynamic. Since the extracted menus can be of various types (e.g. breakfast, beverages, lunch, dinner, etc.), analyze each menu item's description and include only the applicable option blocks. For example, if an item is a beverage or a breakfast item, include only relevant options (like 'Size' or 'Add-ons') and omit blocks that are not applicable. Do not add empty or extraneous option blocks if they are not needed.\n\n"
    "Return only valid JSON that fully conforms to the target structure, without any extra text or formatting."
)

# Function to generate content with retry logic
def generate_content_with_retry(prompt, max_retries=3, retry_delay=5):
    for attempt in range(max_retries):
        try:
            response = client.models.generate_content(
                model=model_id,
                contents=[prompt],
                config={'response_mime_type': 'application/json'}
            )

            if response.parsed:
                return response.parsed

            # Fallback to text parsing if .parsed fails
            if response.text:
                try:
                    return json.loads(response.text)
                except json.JSONDecodeError as e:
                    logging.warning(f"Attempt {attempt + 1}: Failed to parse response.text as JSON: {e}")

            logging.warning(f"Attempt {attempt + 1}: No valid response received.")

        except Exception as e:
            logging.error(f"Attempt {attempt + 1}: An unexpected error occurred: {e}")

        if attempt < max_retries - 1:
            logging.info(f"Retrying in {retry_delay} seconds...")
            time.sleep(retry_delay)

    logging.error("Max retries reached. Failed to generate content.")
    return None

# Generate enhanced menu JSON
enhanced_menu = generate_content_with_retry(prompt)


# Save the enhanced menu JSON
output_directory = "output_files"
os.makedirs(output_directory, exist_ok=True)
output_file_path = os.path.join(output_directory, f"enhanced_menu_{club_name}.json")

if enhanced_menu:
    with open(output_file_path, "w", encoding="utf-8") as outfile:
        json.dump(enhanced_menu, outfile, indent=2)
    print(f"Enhanced Menu saved to: {output_file_path}")
else:
    print("Failed to generate enhanced menu.")
