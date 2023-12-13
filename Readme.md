## Project Overview

This Git repository contains a Golang-based application that provides functionality for user registration, authentication, product management, and order placement. The application utilizes technologies such as Golang, Docker, gRPC, JSON, PostgreSQL, Makefile, and Git.

## Features

1. **User Registration and Authentication:**
   - Users can register and authenticate to access the application.

2. **Product Management:**
   - Users can add products to the system.

3. **Retrieve Products by ID:**
   - Users can retrieve product information using product IDs.

4. **Order Placement:**
   - Users can place orders through the system.

## Technologies Used

- **Golang:** The primary programming language used for development.
- **Docker:** Containerization platform for packaging the application and its dependencies.
- **gRPC:** Remote procedure call framework for communication between services.
- **JSON:** Data interchange format for transmitting data between the client and server.
- **PostgreSQL:** Relational database for storing application data.
- **Makefile:** Used for automation and simplifying build processes.
- **Git:** Version control system for tracking changes in the project.

## Installation

Before proceeding with the installation, ensure that Docker is installed on your system. You can install Docker by following the instructions on the official Docker website: [Install Docker](https://docs.docker.com/get-docker/)

Additionally, Golang and gRPC need to be installed. Please follow the instructions for Golang installation: [Install Golang](https://golang.org/doc/install)

To install gRPC, run the following command:
```bash
go get -u google.golang.org/grpc
```

## Getting Started

Follow these steps to run the application:

1. **Generate gRPC Protobuf Code:**
   ```bash
   make init
   ```


2. **Generate gRPC Protobuf Code:**
   ```bash
   make proto
   ```

3. **Run Tests:**
   ```bash
   make test
   ```

4. **Create PostgreSQL Database:**
   ```bash
   make postgresCreateDB
   ```

5. **Build and Run Docker Container:**
   ```bash
   make dockerRun
   ```

These steps will set up and run the application, making it accessible for use.

## Project Contribution

This project includes an API gateway that processes JSON requests from the client, transforms them into gRPC requests, and forwards them to the appropriate services. Each service interacts with the database to fulfill the client's request.

## Notes

- Ensure that all necessary dependencies are installed and configured before running the application.
- For any issues or improvements, please open an issue or submit a pull request.