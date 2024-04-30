### Graph Structure and Algorithms in Go

#### Project Description:
This project serves as a comprehensive platform for visualizing graph structures and executing various graph algorithms. Currently, it implements essential algorithms such as MST (Minimum Spanning Tree) and Dijkstra's algorithm. However, it's designed to accommodate future expansions, including the incorporation of additional algorithms and features.

The server-side component, developed in Go (Golang), provides robust functionality for graph manipulation and algorithm execution.

The user interface, crafted using HTML, CSS, and JavaScript
#### Usage:

- Before running the program, ensure you have set up your `.env` file with your database password:
  ```
  DB_PASSWORD=YOUR_PASSWORD
  ```

- To run the program:
  ```bash
  make run 
  ```

- To build the program:
  ```bash
  make build
  ```

- To update packages:
  ```bash
  make download_packages
  ```

- To create a new migration:
  ```bash
  make create_migrations
  ```

- To apply migrations to the database:
  ```bash
  make migrate_up POSTGRES_PASSWORD=YOUR_PASSWORD
  ```

- To roll back migrations:
  ```bash
  make migrate_down POSTGRES_PASSWORD=YOUR_PASSWORD
  ```

#### Database Setup:
Ensure that PostgreSQL is installed and running locally. Set up the database connection URL in the `DB_URL` variable in the Makefile. By default, it assumes the database name is `graph_db` and the user is `postgres`. Modify these values if necessary.

#### Server Configuration:
The server configuration is stored in a YAML file. Modify the configuration in `config.yaml` as needed.

#### Screenshots:
*Add screenshots of your program here.*

#### Dependencies:
- Go
- PostgreSQL
- HTML
- CSS
- JavaScript

#### Contributor:
- Andrii Buts

