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
    Handles multiline entries by capturing the content until the start of the next key.
    """
    md_file_url = f"{repo_base_url}/{parameter_id}.md"
    md_content = fetch_file_content(md_file_url)
    data_dict = {}
    current_key = None
    content_buffer = []

    for line in md_content.split('\n'):
        # Check for the start of a new key
        key_start = re.match(r"^(.*?): ```", line)
        if key_start:
            if current_key is not None:
                # Save the previous key with its content, joining multi-line text
                data_dict[current_key] = ''.join(content_buffer).strip('```').strip()
            current_key = key_start.group(1).strip()
            content_buffer = [line.split('```')[1] if '```' in line else '']
        elif current_key:
            # Continue appending lines if within a multi-line section
            content_buffer.append(line)

    # Capture the last entry
    if current_key and content_buffer:
        data_dict[current_key] = ''.join(content_buffer).strip('```').strip()

    return data_dict


def fetch_all_parameters(index_file_url):
    """
    Fetch all parameters from the repository using the index.md file.
    """
    repo_base_url = '/'.join(index_file_url.split('/')[:-1]).replace('blob/', '').replace('github.com',
                                                                                          'raw.githubusercontent.com')
    parameter_ids = extract_parameter_ids(index_file_url)
    all_parameters = {}
    for parameter_id in parameter_ids:
        all_parameters[parameter_id] = fetch_markdown_file_by_id(repo_base_url, parameter_id)
    return all_parameters


# Example usage:
index_file_url = 'https://github.com/The-Grid-Data/TGS/blob/main/tgs-005/doc/index.md'
all_parameters = fetch_all_parameters(index_file_url)
print(json.dumps(all_parameters, indent=4))
