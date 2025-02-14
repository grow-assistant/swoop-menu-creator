# Swoop Menu Creator

A Python-based tool suite for converting PDF menus into structured data and generating database seed files for the Swoop platform.

## Overview

This project automates the process of:
1. Extracting menu data from PDF files.
2. Enhancing the extracted data with additional structured information.
3. Critiquing and improving the menu data for accuracy.
4. Generating Go database seed files for the Swoop platform.

## Components

### 1. Main Script (`main.py`)
The entry point of the application that:
- Loads customer information from `customer_info.json`.
- Presents available clubs for selection.
- Orchestrates the execution of the processing pipeline.
- Creates necessary output directories.
- Manages environment variables for inter-script communication.

### 2. Customer Information (`customer_info.json`)
Stores club and restaurant information including:
- Club names and addresses.
- Restaurant names.
- Paths to PDF menu files.

### 3. Menu Extraction (`extract_menu_from_pdf.py`)
Uses Google's Gemini AI to:
- Process PDF menu files.
- Extract structured menu data.
- Convert unstructured PDF content into a standardized JSON format.
- Save the extracted data for further processing.

### 4. Menu Enhancement (`enhance_menu.py`)
Enriches the extracted menu data by:
- Adding detailed option blocks for items.
- Including customization choices (e.g., cooking preferences, sides).
- Structuring removal options.
- Maintaining consistent pricing information.
- Using Gemini AI for intelligent menu analysis.

### 5. Menu Critique (`critique_menu.py`)
Validates and improves menu accuracy by:
- Comparing the PDF, extracted, and enhanced menu versions.
- Identifying and correcting discrepancies.
- Verifying prices and descriptions.
- Ensuring proper option configurations.
- Using Gemini AI for comprehensive review.

### 6. Seed File Creation (`create_seed_file.py`)
Generates Go language seed files that:
- Follow the Swoop platform's database schema.
- Create properly formatted database entries.
- Handle location, menu, category, and item relationships.
- Include all menu options and customizations.

## Setup

1. Install required dependencies:
   ```bash
   pip install -r requirements.txt
   ```

2. Create a `.env` file with your Google Gemini API key:
   ```plaintext
   GOOGLE_GEMINI_API=your_api_key_here
   ```

3. Place PDF menu files in the `pdf_files` directory.

4. Update `customer_info.json` with club information.

## Usage

Run the main script:
   ```bash
   python main.py
   ```

Follow the prompts to:
1. Select a club to process
2. Wait for the extraction and enhancement process
3. Review the generated files in:
   - `output_files/` for JSON data
   - `seed_files/` for Go seed files

## Output Files

The process generates several files:
- `extracted_menu_{club_name}.json`: Raw extracted menu data
- `enhanced_menu_{club_name}.json`: Enriched menu structure
- `improved_enhanced_menu_{club_name}.json`: Final, critiqued menu data
- `{club_name}.go`: Database seed file for the Swoop platform

## Dependencies

- Python 3.8+
- Google Gemini AI API
- Pydantic for data validation
- python-dotenv for environment management

## Notes

- Ensure PDF files are text-searchable for optimal extraction.
- The Gemini AI model requires an active internet connection.
- Generated seed files should be reviewed before database insertion.
