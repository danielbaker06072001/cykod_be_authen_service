# This project is implemented based on CleanArchitecture , designed by BOB

    ! Document

[https://medium.com/@jamal.kaksouri/building-better-go-applications-with-clean-architecture-a-practical-guide-for-beginners-98ea061bf81a]

    ! Video

[https://www.youtube.com/watch?v=ffYCgcDgsfw]

# How to create migrate file

migrate create -ext sql -dir Migration/sql -seq create_users_table

# How to run migrate up

make <command_in_makefile>

# How to connect to redis from docker, and how to check if bit exist

docker exec -it authentication_redis redis-cli

-   CHECK BIT: getbit u:bit 4097051691
-   SET BIT : setbit u:bit 4097051691

# Authentication Services

This project implements authentication services for user management, including functionalities for user registration, login, and password hashing with salt. The services are written in Go and use a secure hashing algorithm to protect user passwords.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Overview

This project provides a robust authentication system to manage user registration and login processes. It securely stores user passwords using a combination of hashing and salting techniques to ensure security.

## Features

- User Registration
- User Login
- Password Hashing with Salt
- User Authentication
- Error Handling

## Getting Started

### Prerequisites

- Go 1.17 or later
- A running instance of a SQL database (e.g., MySQL, PostgreSQL)
- Git

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/authentication-services.git
   cd authentication-services
   ```

2. **Install dependencies:**
    ```bash
    go mod tidy
    ```

## Configuration
- *Database Configuration*: Configure your database connection in the config package. Update the config.go file with your database connection details.

- *Environment Variables*: Create a .env file in the root directory of the project and set the following variables:

makefile
```bash
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
```

**Usage**
1. Run the service:
```bash
go run main.go
Register a new user: Send a POST request to /register with the following JSON body:
```

2. Registering a new User: Send a POST request to /register with the following JSOn
```bash
{
    "username": "test",
    "password": "Test",
    "first_name": "test",
    "last_name": "test",
    "email": "test1@gmail.com"
}
```

3. Login with an existing user: Send a POST request to /login with the following JSON body:
```bash
{
  "username": "exampleuser",
  "password": "examplepassword"
}
```

**API endpoints**
POST /register: Registers a new user.
POST /login: Authenticates a user and returns a success message if the credentials are valid.

**Example Request**
1. Register:
```bash
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "exampleuser", "password": "examplepassword"}'
```
2. Login:
```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "exampleuser", "password": "examplepassword"}'
```

