import pandas as pd
import mysql.connector
from mysql.connector import Error

from credentials import db_config

# Path to the CSV file
csv_file_path = 'tgs.csv'

# Function to load data from CSV and insert into the database
def load_data(csv_path):
    # Load data from CSV
    data = pd.read_csv(csv_path)

    # Connect to the MySQL database
    try:
        conn = mysql.connector.connect(**db_config)
        if conn.is_connected():
            print('Connected to MySQL database')

            cursor = conn.cursor()

            # Disable autocommit for transaction
            conn.autocommit = False

            # Create a list of tuples from the dataframe values
            tuples = [tuple(x) for x in data.to_numpy()]

            # Comma-separated dataframe columns, adjusting for SQL-compatible names
            cols = ', '.join('`{0}`'.format(col.replace(" ", "_").replace(":", "")) for col in data.columns)  # Handling space in column names

            # Constructing the query string
            query = "INSERT INTO DbTGS ({cols}) VALUES (%s{placeholders})".format(
                cols=cols,
                placeholders=', %s' * (len(data.columns) - 1))

            # Execute the query as a single transaction
            cursor.executemany(query, tuples)
            conn.commit()
            print("Data inserted successfully")
        else:
            print("Connection failed")
    except Error as e:
        print(f"Error: {e}")
        # Rollback in case of error
        conn.rollback()
    finally:
        if conn.is_connected():
            cursor.close()
            conn.close()
            print('Connection closed.')

# Execute the function
load_data(csv_file_path)
