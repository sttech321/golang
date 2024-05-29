# golang
GoLang crud opertaion

# User CRUD Operation

This project is a sample Go application that performs CRUD (Create, Read, Update, Delete) operations for managing users.

## Table of Contents

- [User CRUD Operation](#user-crud-operation)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Steps](#steps)
  - [Usage](#usage)
    - [Running the application](#running-the-application)
    - [Building the application](#building-the-application)
    - [Running tests](#running-tests)
  - [API Endpoints](#api-endpoints)
  - [Features](#features)
  - [Configuration](#configuration)
    - [Environment Variables](#environment-variables)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

### Prerequisites

- Go version 1.XX or higher
- Any other software or tools needed

### Steps

1. Clone the repository: 
   git clone https://github.com/sttech321/golang.git



Installation
  
Install dependencies:
 
go mod tidy
 
 
To run the application, execute:
 
go run main.go 

To build the application, execute:
 
go build -o usercrud main.go


Running tests
To run tests, execute:

sh
Copy code
go test ./...
API Endpoints
POST /users: Create a new user
GET /users: Get a list of all users
GET /users/{id}: Get a single user by ID
PUT /users/{id}: Update a user by ID
DELETE /users/{id}: Delete a user by ID
Features
Create new users
Retrieve a list of users
Retrieve a specific user by ID
Update user details
Delete a user
