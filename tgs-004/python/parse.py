import json
import os
import re

def parse_markdown_file(filepath):
    """
    Parse the content of a Markdown file and convert it into a dictionary.
    """
    with open(filepath, 'r') as file:
        content = file.read()

    # Extract the main parameter ID from the first line (used as the title in the format: "# ID")
    parameter_id = content.split('\n')[0].strip('# ').strip()

    # Dictionary to store the data extracted from the Markdown
    data_dict = {'Parameter ID': parameter_id}

    # Extract key-value pairs from the content
    content_lines = content.split('\n')
    for line in content_lines:
        match = re.match(r"^(.*?): ```(.*?)```$", line)
        if match:
            key = match.group(1).strip()
            value = match.group(2).strip()
            data_dict[key] = value

    return data_dict

def read_markdown_files_to_dict(folder_path):
    """
    Read all Markdown files in the specified folder and return a dictionary of dictionaries,
    where each key is a Parameter ID and its value is the corresponding data dictionary.
    """
    markdown_files = [f for f in os.listdir(folder_path) if f.endswith('.md') and f != 'index.md']

    all_data = {}
    for md_file in markdown_files:
        filepath = os.path.join(folder_path, md_file)
        file_data = parse_markdown_file(filepath)
        param_id = file_data.get('Parameter ID')
        if param_id:
            all_data[param_id] = file_data

    return all_data

# Example usage:
# Replace '../doc/' with the actual path where your Markdown files are stored.
folder_path = '../doc/'
markdown_data_dict = read_markdown_files_to_dict(folder_path)
#print(markdown_data_dict)
print(json.dumps(markdown_data_dict, indent=4))