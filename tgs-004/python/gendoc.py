import os
import pandas as pd
import numpy as np


def generate_markdown_content(row):
    """
    Generate Markdown content based on a row from the DataFrame.
    """
    # Extracting individual fields and replacing NaN values with empty strings
    name = row.get('Name', '') if not pd.isna(row.get('Name', '')) else ''
    parameter_id = row.get('Parameter ID', '') if not pd.isna(row.get('Parameter ID', '')) else ''
    type_ = row.get('Type', '') if not pd.isna(row.get('Type', '')) else ''
    specs = row.get('Specifications', '') if not pd.isna(row.get('Specifications', '')) else ''
    description = row.get('Description', '') if not pd.isna(row.get('Description', '')) else ''
    validation_steps = row.get('Validation steps', '') if not pd.isna(row.get('Validation steps', '')) else ''
    in_dbd = row.get('In DBD', '') if not pd.isna(row.get('In DBD', '')) else ''
    notes = row.get('Notes:', '') if not pd.isna(row.get('Notes:', '')) else ''

    # Creating the Markdown content
    md_content = f"# {parameter_id}\n\n"

    for col, val in row.items():
        if pd.isna(val):
            val = ''  # Replace NaN with an empty string
        md_content += f"{col}: {val}\n\n"

    return md_content


def generate_markdown_files(csv_file, target_folder):
    # Reading the CSV file into a DataFrame
    df = pd.read_csv(csv_file)

    # Ensuring the target folder exists
    if not os.path.exists(target_folder):
        os.makedirs(target_folder)

    # Iterating over each row in the DataFrame
    for _, row in df.iterrows():
        # Extracting the parameter ID
        parameter_id = row['Parameter ID']

        # Generating the Markdown content for this row
        md_content = generate_markdown_content(row)

        # Creating the file path
        file_path = os.path.join(target_folder, f"{parameter_id}.md")

        # Writing the Markdown content to the file
        with open(file_path, 'w') as f:
            f.write(md_content)

    print(f"Markdown files have been generated in {target_folder}")


# Usage example:
# Make sure to replace 'tgs.csv' with the path to your actual CSV file.
csv_file = 'tgs.csv'
target_folder = 'generated_md_files'
generate_markdown_files(csv_file, target_folder)
