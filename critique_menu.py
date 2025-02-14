import os
import json
from google import genai
from dotenv import load_dotenv
import logging
import time

# Setup logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# Load API key and environment variables
load_dotenv('.env')
api_key = os.getenv("GOOGLE_GEMINI_API")
client = genai.Client(api_key=api_key)
model_id = "gemini-2.0-flash"

def load_json_file(filepath):
    try:
        with open(filepath, "r", encoding="utf-8") as infile:
            return infile.read()
    except Exception as e:
        logging.error(f"Error reading {filepath}: {e}")
        exit(1)

# Get the club name either from environment or use default
club_name = os.getenv("CLUB_NAME", "gates_four_golf_&_country_club")
club_name_normalized = club_name.replace(" ", "_").lower()

# Load the enhanced menu JSON from file using the club-specific filename
enhanced_menu_file = os.path.join("output_files", f"enhanced_menu_{club_name_normalized}.json")
extracted_menu_file = os.path.join("output_files", f"extracted_menu_{club_name_normalized}.json")

# Load both JSON files
enhanced_menu_json = load_json_file(enhanced_menu_file)
extracted_menu_json = load_json_file(extracted_menu_file)

# Retrieve the PDF file path from environment variable or use default
pdf_path = os.getenv("PDF_PATH", "pdf_files/gatesFour.pdf")

# Ensure the PDF file exists
if not os.path.exists(pdf_path):
    logging.error(f"PDF file not found at path: {pdf_path}")
    exit(1)

# Upload the PDF file
logging.info(f"Processing PDF file: {pdf_path}")
pdf_file = client.files.upload(file=pdf_path, config={'display_name': club_name})

# Construct the critique prompt
critique_prompt = (
    "Compare these three sources of menu information and identify any discrepancies or missing details:\n"
    "1. The original PDF menu (attached)\n"
    "2. The extracted menu JSON:\n"
    f"{extracted_menu_json}\n\n"
    "3. The enhanced menu JSON:\n"
    f"{enhanced_menu_json}\n\n"
    "Please analyze and improve the enhanced menu JSON by:\n"
    "1. Correcting any incorrect prices or descriptions, while maintaining the existing structure\n"
    "2. Ensuring all 'Remove Options' blocks have min=0 and max=0\n"
    "3. Verifying that required choices (like meat temperature, sides, etc.) are properly specified\n"
    "4. DO NOT restructure or reorganize existing option blocks - only correct prices and descriptions\n"
    "5. DO NOT combine or split existing menu items or their options\n\n"
    "Important: Preserve the exact structure of all existing items and their option blocks. \n"
    "Do not create new option blocks or reorganize existing ones. \n"
    "Focus only on correcting prices, descriptions, and min/max values.\n\n"
    "Return only the improved enhanced menu JSON with no additional commentary."
)

# Define a helper function for retrying the content generation
def generate_content_with_retry(prompt, contents, max_retries=3, retry_delay=5):
    for attempt in range(max_retries):
        try:
            response = client.models.generate_content(
                model=model_id,
                contents=contents,
                config={'response_mime_type': 'application/json'}
            )
            if response.parsed:
                return response.parsed
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

# Build the contents for the critique
contents = [critique_prompt, pdf_file]

# Generate the improved enhanced menu JSON
improved_menu = generate_content_with_retry(critique_prompt, contents)

# Save the improved menu JSON
output_directory = "output_files"
os.makedirs(output_directory, exist_ok=True)
improved_menu_file = os.path.join(output_directory, f"improved_enhanced_menu_{club_name_normalized}.json")

if improved_menu:
    with open(improved_menu_file, "w", encoding="utf-8") as outfile:
        json.dump(improved_menu, outfile, indent=2)
    print(f"Improved Enhanced Menu saved to: {improved_menu_file}")
else:
    print("Failed to generate improved enhanced menu.")

