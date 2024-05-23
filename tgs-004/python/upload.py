import pandas as pd
import mysql.connector
from mysql.connector import Error

from credentials import db_config

# Path to the CSV file
csv_file_path = 'tgs.csv'

# Function to load data from CSV and insert/update into the database
def load_data(csv_path):
    # Load data from CSV
    data = pd.read_csv(csv_path)

    # Making sure data columns match the database columns, including renaming if necessary
    data.columns = data.columns.str.replace(" ", "_").str.replace(":", "")  # Adjust column names to match DB columns

    # Check for columns that need type conversion, particularly boolean/tinyint types
    if 'In_DBD' in data.columns:
        data['In_DBD'] = data['In_DBD'].astype(int)

    # Connect to the MySQL database
    try:
        with mysql.connector.connect(**db_config) as conn:
            if conn.is_connected():
                print('Connected to MySQL database')

                with conn.cursor() as cursor:
                    # Disable autocommit for transaction
                    conn.autocommit = False

                    # Create a list of tuples from the dataframe values
                    tuples = [tuple(x) for x in data.to_numpy()]

                    # Comma-separated dataframe columns
                    cols = ', '.join(f"`{col}`" for col in data.columns)

                    # Constructing the query string with ON DUPLICATE KEY UPDATE clause
                    update_cols = ', '.join([f"`{col}` = VALUES(`{col}`)" for col in data.columns if col != 'Parameter_ID'])
                    query = f"INSERT INTO `DbTGS` ({cols}) VALUES ({', '.join(['%s'] * len(data.columns))}) ON DUPLICATE KEY UPDATE {update_cols}"

                    # Execute the query as a single transaction
                    cursor.executemany(query, tuples)
                    conn.commit()
                    print("Data inserted/updated successfully")
    except Error as e:
        print(f"Error: {e}")
        if conn.is_connected():
            conn.rollback()
    finally:
        # The connection is auto-closed by the 'with' statement, but we ensure it if needed
        if conn.is_connected():
            print('Connection closed.')

# Execute the function
load_data(csv_file_path)
