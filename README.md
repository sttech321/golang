# golang
GoLang crud opertaion


Installation
Prerequisites
Go version 1.XX or higher
Any other software or tools needed
Steps
Clone the repository:
sh
Copy code
git clone https://github.com/sttech321/golang.git
Navigate to the project directory:
sh
Copy code
cd user-crud-sample
Install dependencies:
sh
Copy code
go mod tidy
Usage
Running the application
To run the application, execute:

sh
Copy code
go run main.go
Building the application
To build the application, execute:

sh
Copy code
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
