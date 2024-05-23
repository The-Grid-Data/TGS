import requests
import re
import json

def fetch_file_content(url):
    """
    Fetch the content of a file from a given URL.
    """
    response = requests.get(url)
    response.raise_for_status()  # Raises HTTPError for bad requests (4XX, 5XX)
    return response.text

def extract_parameter_ids(index_file_url):
    """
    Read the index.md from a GitHub repository URL and return a list of parameter IDs.
    """
    raw_index_url = index_file_url.replace('github.com', 'raw.githubusercontent.com').replace('blob/', '')
    index_content = fetch_file_content(raw_index_url)
    parameter_ids = re.findall(r"\((.*?).md\)", index_content)
    return parameter_ids

def fetch_markdown_file_by_id(repo_base_url, parameter_id):
    """
    Fetch and parse a Markdown file from a GitHub repository by parameter ID.
    """
    md_file_url = f"{repo_base_url}/{parameter_id}.md"
    md_content = fetch_file_content(md_file_url)
    data_dict = {'Parameter ID': parameter_id}
    lines = md_content.split('\n')
    for line in lines:
        match = re.match(r"^(.*?): ```(.*?)```$", line)
        if match:
            key = match.group(1).strip()
            value = match.group(2).strip()
            data_dict[key] = value
    return data_dict

def fetch_all_parameters(index_file_url):
    """
    Fetch all parameters from the repository using the index.md file.
    """
    repo_base_url = '/'.join(index_file_url.split('/')[:-1]).replace('blob/', '').replace('github.com', 'raw.githubusercontent.com')
    parameter_ids = extract_parameter_ids(index_file_url)
    all_parameters = {}
    for parameter_id in parameter_ids:
        all_parameters[parameter_id] = fetch_markdown_file_by_id(repo_base_url, parameter_id)
    return all_parameters

# Example usage:
index_file_url = 'https://github.com/The-Grid-Data/TGS/blob/main/tgs-005/doc/index.md'
repo_base_url = '/'.join(index_file_url.split('/')[:-1]).replace('blob/', '').replace('github.com',
                                                                                      'raw.githubusercontent.com')

all_parameters = fetch_all_parameters(index_file_url)
print(json.dumps(all_parameters, indent=4))
print(repo_base_url)