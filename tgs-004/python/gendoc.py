import os
import pandas as pd


def generate_markdown_files(csv_file, target_folder):
    # Reading the CSV file into a DataFrame
    df = pd.read_csv(csv_file)

    # Ensuring the target folder exists
    if not os.path.exists(target_folder):
        os.makedirs(target_folder)

    # Iterating over each row in the DataFrame
    for _, row in df.iterrows():
        # Extracting necessary information from the row
        parameter_id = row['Parameter ID']
        name = row.get('Name', '')
        description = row.get('Description', '')
        specs = row.get('Specifications', '')
        validation_steps = row.get('Validation steps', '')
        notes = row.get('Notes:', '')

        # Generating the content for the Markdown file
        md_content = f"# {name}\n\n"

        if description:
            md_content += f"## Description\n{description}\n\n"

        if specs:
            md_content += f"## Specifications\n{specs}\n\n"

        if validation_steps:
            md_content += f"## Validation Steps\n{validation_steps}\n\n"

        if notes:
            md_content += f"## Notes\n{notes}\n\n"

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
