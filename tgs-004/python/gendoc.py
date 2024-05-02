import os
from copy import Error

import pandas as pd


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

    generated_files_group = []
    generated_files_dict = {}
    # Iterating over each row in the DataFrame
    for _, row in df.iterrows():
        # Extracting the parameter ID
        parameter_id = row['Parameter ID']
        if f"{parameter_id}" == "nan":
            continue

            # Splitting the Parameter ID by "."
        param_tokens = parameter_id.split('.') if parameter_id else []
        if len(param_tokens) >= 2:
            param_prefix, param_suffix = param_tokens[0], param_tokens[1]
        else:
            raise Error(f"{parameter_id}")

        if param_prefix not in generated_files_dict:
            generated_files_group.append(param_prefix)
            generated_files_dict[param_prefix] = []

        generated_files_dict[param_prefix].append(param_suffix)

        # Generating the Markdown content for this row
        md_content = generate_markdown_content(row)

        # Creating the file path
        file_path = os.path.join(target_folder, f"{parameter_id}.md")

        if f"{parameter_id}" == "nan":
            break

        # Writing the Markdown content to the file
        with open(file_path, 'w') as f:
            f.write(md_content)

    print(f"Markdown files have been generated in {target_folder}")

    # Writing the final summary file
    summary_file_path = os.path.join(target_folder, "index.md")
    with open(summary_file_path, 'w') as summary_file:
        summary_file.write("# Index\n\n")
        for pr in generated_files_group:
            summary_file.write(f"# {pr}\n\n")
            for po in generated_files_dict[pr]:
                # Extracting just the filename
                filename = f"{pr}.{po}"

                summary_file.write(f"- [{filename}]({filename}.md)\n")

    print(f"Markdown files and summary have been generated in {target_folder}")


# Usage example:
# Make sure to replace 'tgs.csv' with the path to your actual CSV file.
csv_file = 'tgs.csv'
target_folder = '../doc/'
generate_markdown_files(csv_file, target_folder)
