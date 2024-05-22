# Graph Project

## Overview

This project focuses on graph structures and algorithms such as Dijkstra's algorithm and Minimum Spanning Tree (MST). Our main server is built using Go with the Gin framework, the frontend is developed using JavaScript and the vis.js library, and we also have a graph recognition module from images using Python with OpenCV and NumPy libraries. The graph recognition server is hosted using Flask.

## Project Structure

- **Backend (Go):**
  - **Framework:** Gin
  - **Main File:** `cmd/main.go`
- **Frontend (JavaScript):**
  - **Library:** vis.js
- **Graph Recognition (Python):**
  - **Libraries:** OpenCV, NumPy
  - **Server:** Flask
  - **Main File:** `graph-recognition/main.py`

## Prerequisites

- Go installed
- Python 3 installed
- PostgreSQL installed and running
- Make installed

## Setup Instructions

### 1. Clone the Repository

```sh
git clone <repository_url>
cd <repository_name>
```

### 2. Setting Up Environment Variables

To set up the necessary environment variables, follow these steps:

1. **Modify `.env` File:**
   Open the copied `.env` file and replace `YOUR_PASSWORD` with your actual PostgreSQL password. Your `.env` file should look like this:
   ```plaintext
   DB_PASSWORD=YOUR_PASSWORD
   ```
   Save the changes.

2. **Configure `config.yaml` File:**
   Navigate to the `configs` directory and locate the `config.yaml` file. Open it in a text editor.
   You can adjust the database connection settings according to your environment. Here's an example of the configuration section:
   ```yaml
   db:
     host: "localhost"
     port: "5432"
     user: "postgres"
     db_name: "graph_db"
   ```
   Modify the values as needed. Ensure that the `host`, `port`, `user`, and `db_name` fields match your PostgreSQL configuration.
   Additionally, you can adjust the `bind_addr` value if you need to change the server's binding address and port.
   Save the changes to `config.yaml`.


### 3. Combined Commands

- **Run All Go Commands:**

  ```sh
  make all_go
  ```

- **Run All Python Commands (for macOS/Linux):**

  ```sh
  make all_python_mac
  ```

- **Run All Python Commands (for Windows):**

  ```sh
  make all_python_windows
  ```

### 4. Backend Setup (Go)

1. **Download Packages:**

   ```sh
   make download_packages
   ```

2. **Build the Project:**

   ```sh
   make build
   ```

3. **Run the Server:**

   ```sh
   make run
   ```

4. **Database Migrations:**

  - Create Migrations:
    ```sh
    make create_migrations
    ```
  - Apply Migrations:
    ```sh
    make migrate_up POSTGRES_PASSWORD=YOUR_PASSWORD
    ```
  - Rollback Migrations:
    ```sh
    make migrate_down POSTGRES_PASSWORD=YOUR_PASSWORD
    ```


### 5. Graph Recognition Setup (Python)

1. **Create Virtual Environment:**

   ```sh
   make create_venv
   ```

2. **Activate Virtual Environment:**

  - For Windows:
    ```sh
    make activate_venv_windows
    ```
  - For macOS/Linux:
    ```sh
    make activate_venv_macOs
    ```

3. **Install Requirements:**

   ```sh
   make install_requirements
   ```

4. **Run Flask Server:**

   ```sh
   make run_flask
   ```


## Contributors

* [Vanya Shumylo](https://github.com/VanyaShumilo1)
* [Andrii Buts](https://github.com/buts00)
* [Mykola Hreba](https://github.com/heckq)



