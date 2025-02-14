from google import genai
from pydantic import BaseModel, Field, PydanticUserError
import os
from dotenv import load_dotenv
import logging

load_dotenv('.env')
api_key = os.getenv("GOOGLE_GEMINI_API") # If you are not using Colab you can set the API key directly

# Create a client
client = genai.Client(api_key=api_key)

# Define the model you are going to use
model_id =  "gemini-2.0-flash" # or "gemini-2.0-flash-lite-preview-02-05"  , "gemini-2.0-pro-exp-02-05"
     
#Gemini models are able to process images and videos, which can used with base64 strings or using the files api.
#The Python API includes a upload and delete method.
def extract_structured_data(file_path: str, model: BaseModel):
    # Upload the file to the File API
    file = client.files.upload(file=file_path, config={'display_name': file_path.split('/')[-1].split('.')[0]})
    # Generate a structured response using the Gemini API
    prompt = (
        "Extract the structured menu data from the provided food menu. "
        "Return the JSON following this structure: "
        "{ 'menu': { "
        "    'name': string, "
        "    'location': string, "
        "    'categories': [ { "
        "         'name': string, "
        "         'description': string, "
        "         'items': [ { "
        "              'name': string, "
        "              'description': string, "
        "              'price': number, "
        "              'options': [ { 'name': string, 'price': number } ] "
        "         } ] "
        "    } ] "
        "} }"
    )
    response = client.models.generate_content(model=model_id, contents=[prompt, file], config={'response_mime_type': 'application/json', 'response_schema': model})
    # Convert the response to the pydantic model and return it
    return response.parsed

class MenuItemOption(BaseModel):
    name: str = Field(description="The name of the menu item option")
    price: float = Field(description="The price for the option")

class MenuItem(BaseModel):
    name: str = Field(description="The name of the menu item")
    description: str = Field(description="The description of the menu item")
    price: float = Field(description="The base price of the menu item")
    options: list[MenuItemOption] = Field(default_factory=list, description="List of menu item options")
  
class MenuCategory(BaseModel):
    name: str = Field(description="The name of the category")
    items: list[MenuItem] = Field(default_factory=list, description="List of menu items in this category")

class Menu(BaseModel):
    name: str = Field(description="The name of the menu")
    location: str = Field(description="The location identifier for the menu")
    categories: list[MenuCategory] = Field(default_factory=list, description="List of menu categories")

class MenuStructure(BaseModel):
    menu: Menu

# Provide the path to your PDF file from environment variable (set by main.py)
pdf_path = os.getenv("PDF_PATH", "pdf_files/gatesFour.pdf")

# Ensure the PDF file exists before trying to process it
if not os.path.exists(pdf_path):
    logging.error(f"PDF file not found at path: {pdf_path}")
    exit(1)

logging.info(f"Processing PDF file: {pdf_path}")

# Extract the structured data
result = extract_structured_data(pdf_path, MenuStructure)

# Ensure the output directory exists
output_directory = "output_files"
os.makedirs(output_directory, exist_ok=True)

# Use the club name (normalized) for the filename
club_name = os.getenv("CLUB_NAME", "gatesFour").replace(" ", "_").lower()
output_file_path = os.path.join(output_directory, f"extracted_menu_{club_name}.json")

# Save the result to a file
with open(output_file_path, "w", encoding="utf-8") as output_file:
    output_file.write(result.model_dump_json())

print(f"Extracted Menu saved to: {output_file_path}")

