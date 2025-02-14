import json
import os
import subprocess
import sys

def main():
    # Load customer info from customer_info.json
    customer_info_file = "customer_info.json"
    try:
        with open(customer_info_file, "r", encoding="utf-8") as f:
            customer_data = json.load(f)
    except Exception as e:
        print(f"Error reading {customer_info_file}: {e}")
        sys.exit(1)
    
    clubs = customer_data.get("clubs", [])
    if not clubs:
        print("No clubs found in customer info.")
        sys.exit(1)
    
    # List available clubs for selection
    print("Available Clubs:")
    for idx, club in enumerate(clubs, 1):
        print(f"{idx}. {club['name']} - {club['address']} ({club['restaurant']})")
    
    # Automatically select Gates Four (index 0)
    choice_index = 0
    print("Automatically selecting Gates Four Golf & Country Club")
    
    selected_club = clubs[choice_index]
    pdf_path = selected_club["pdfMenuPath"]
    club_name = selected_club["name"]
    
    # Convert the path to use proper system separators
    pdf_path = os.path.normpath(pdf_path)
    
    # Verify PDF file exists
    if not os.path.exists(pdf_path):
        print(f"Error: PDF file not found at path: {pdf_path}")
        sys.exit(1)
    
    print(f"Selected Club: {club_name}")
    print(f"Using PDF Menu Path: {pdf_path}")
    
    # Set environment variables for the other scripts to use
    os.environ["PDF_PATH"] = pdf_path
    os.environ["CLUB_NAME"] = club_name
    
    # Create required directories if they don't exist
    output_directory = "output_files"
    pdf_directory = "pdf_files"
    os.makedirs(output_directory, exist_ok=True)
    os.makedirs(pdf_directory, exist_ok=True)
    
    # Run extract_menu_from_pdf.py to generate the extracted menu JSON
    print("Running extract_menu_from_pdf.py ...")
    result = subprocess.run(
        ["python", "extract_menu_from_pdf.py"],
        capture_output=True,
        text=True
    )
    if result.returncode != 0:
        print("Error running extract_menu_from_pdf.py:")
        print(result.stderr)
        sys.exit(1)
    else:
        print(result.stdout)
    
    # At this point, the file "output_files/extracted_menu_{club_name}.json" should be generated.
    # Now run enhance_menu.py to process the extracted JSON
    print("Running enhance_menu.py ...")
    result = subprocess.run(
        ["python", "enhance_menu.py"],
        capture_output=True,
        text=True
    )
    if result.returncode != 0:
        print("Error running enhance_menu.py:")
        print(result.stderr)
        sys.exit(1)
    else:
        print(result.stdout)
    
    # New addition: Run critique_menu.py
    print("Running critique_menu.py ...")
    result = subprocess.run(
        ["python", "critique_menu.py"],
        capture_output=True,
        text=True
    )
    if result.returncode != 0:
        print("Error running critique_menu.py:")
        print(result.stderr)
        sys.exit(1)
    else:
        print(result.stdout)
    
    # New addition: Run create_seed_file.py
    print("Running create_seed_file.py ...")
    result = subprocess.run(
        ["python", "create_seed_file.py"],
        capture_output=True,
        text=True
    )
    if result.returncode != 0:
        print("Error running create_seed_file.py:")
        print(result.stderr)
        sys.exit(1)
    else:
        print(result.stdout)
    
    print("Process completed successfully.")

if __name__ == "__main__":
    main()   