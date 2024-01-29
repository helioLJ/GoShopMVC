# Product CRUD Application in Golang

## Overview

This is a simple CRUD (Create, Read, Update, Delete) application written in Golang, following the MVC (Model-View-Controller) architecture. The application allows you to manage products, providing basic functionalities to create, retrieve, update, and delete product information.

## Features

- **Create:** Add new products with details such as name, description, and price.
- **Read:** Retrieve a list of all products or get details of a specific product.
- **Update:** Modify product information, including name, description, and price.
- **Delete:** Remove products from the system.

## Technologies Used

- Golang: The primary programming language.
- pq: Pure Go Postgres driver for database/sql.
- Postgres: The open-source relational database management system.
- Docker: A platform for developing, shipping, and running applications using containerization.

## Project Structure
The project structure consists of the following directories and files:

```
├── controllers
│   └── produtos.go
├── db
│   └── db.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── models
│   └── produtos.go
├── README.md
├── routes
│   └── routes.go
└── templates
    ├── edit.html
    ├── _head.html
    ├── index.html
    ├── _menu.html
    └── new.html

5 directories, 14 files
```
- **controllers**: This folder contains the controller files for handling different aspects of the application's logic. In this case, it includes a produtos.go file, which likely contains the controller logic for managing products.

- **db**: This folder contains the database-related files. The db.go file is likely responsible for establishing a connection to the database and providing functions for interacting with it.

- **models**: This folder contains the model files that define the structure and behavior of the application's data. In this case, it includes a produtos.go file, which likely defines the product model.

- **routes**: This folder contains the route files that define the API endpoints and their corresponding handlers. In this case, it includes a routes.go file, which likely defines the routes for managing products.

- **templates**: This folder contains the HTML templates used for rendering the views of the application. It includes various HTML files such as edit.html, index.html, new.html, etc., which define the structure and content of different pages.

The project structure is organized in a way that separates concerns and promotes modularity. This allows for easier maintenance and scalability of the application.

## Getting Started

1. Ensure you have Golang installed on your machine.
2. Clone this repository: `git clone https://github.com/helioLJ/GoShopMVC.git`
3. Change into the project directory: `cd GoShopMVC`
4. Run the application: `go run main.go`

## API Endpoints

- `GET /`: This endpoint displays the index page of the application.
- `GET /new`: This endpoint displays the page for creating a new item.
- `POST /insert`: This endpoint is used to insert a new item into the database.
- `POST /delete`: This endpoint is used to delete an existing item from the database.
- `GET /edit`: This endpoint displays the page for editing an existing item.
- `POST /update`: This endpoint is used to update an existing item in the database.

## Dependencies

- [pq](https://github.com/lib/pq): Pure Go Postgres driver for database/sql.

## Contributing

Feel free to contribute by submitting issues or pull requests. Your feedback and collaboration are welcome!
