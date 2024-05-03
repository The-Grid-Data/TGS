import json
import requests
import re

def fetch_file_content(url):
    """
    Fetch the content of a file from a given URL.
    """
    response = requests.get(url)
    response.raise_for_status()  # Will raise an HTTPError for bad requests (4XX, 5XX)
    return response.text

def parse_markdown_content(content):
    """
    Parse the content of a Markdown file and convert it into a dictionary.
    """
    parameter_id = content.split('\n')[0].strip('# ').strip()
    data_dict = {'Parameter ID': parameter_id}

    content_lines = content.split('\n')
    for line in content_lines:
        match = re.match(r"^(.*?): ```(.*?)```$", line)
        if match:
            key = match.group(1).strip()
            value = match.group(2).strip()
            data_dict[key] = value

    return data_dict

def read_markdown_files_from_github(repo_base_url, index_file_url):
    """
    Read all Markdown files listed in the index.md from a GitHub repository and return a dictionary of dictionaries.
    """
    # Fetch the index.md content
    index_content = fetch_file_content(index_file_url)
    all_data = {}

    # Extract Markdown file links from the index content
    lines = index_content.split('\n')
    for line in lines:
        match = re.search(r"\((.*?)\.md\)", line)
        if match:
            md_file_url = f"{repo_base_url}/{match.group(1)}.md"
            md_content = fetch_file_content(md_file_url)
            file_data = parse_markdown_content(md_content)
            param_id = file_data.get('Parameter ID')
            if param_id:
                all_data[param_id] = file_data

    return all_data

# Example usage:
repo_base_url = 'https://raw.githubusercontent.com/The-Grid-Data/TGS/main/tgs-004/doc'
index_file_url = 'https://github.com/The-Grid-Data/TGS/blob/main/tgs-004/doc/index.md'.replace('github.com', 'raw.githubusercontent.com').replace('blob/', '')
markdown_data_dict = read_markdown_files_from_github(repo_base_url, index_file_url)
print(json.dumps(markdown_data_dict, indent=4))
