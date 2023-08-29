# User Search CLI

This is a guide on how to run the Golang project. The project was developed using Golang version 1.21.0.

## Prerequisites

Before you start, ensure you have the following installed on your machine:

- Golang 1.21.0: You can download and install it from the [official Golang website](https://go.dev/dl/).


## Configuration

The project uses configurations from an existing `.env` file. Ensure that the following configurations are correctly set in the `.env` file:

- **API_URLS:** This configuration should hold a list of API URLs separated by commas.
- **CSV_FILE:** This configuration should specify the path to the CSV file used by the application.


## Getting Started

Follow these steps to get the project up and running:

1. **Extract the Zip File:**
   Extract the contents of the downloaded zip file to a directory of your choice.

2. **Navigate to the Project Directory:**
   Open a terminal and navigate to the directory where you extracted the zip file:
3. **Build the Project:**
   Run the following command to build the project:
    ```makefile
    make build
    ```
   This will create an executable file with the project name in the same directory.

4. **Run the Project:**
   Run the following command to fetch and search user data by tags:
    ```makefile
    make run tags=sometag,anothertag
    `````
   This will create csv file in the same directory. When a user is found with given tags, its name and salary should be written on the output as a list

5. **Run the Project without fetch:**
   Run the following command to search user data by tags:
    ```makefile
    make run tags=sometag,anothertag fetch=false
    `````
    When a user is found with given tags, its name and salary should be written on the output as a list

6. **Run the tests:**
   Run the following command get test coverage:
    ```makefile
    make test
    `````