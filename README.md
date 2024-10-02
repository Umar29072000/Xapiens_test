# Xapiens_test

## Overview

**Xapiens_test** is a Go-based application designed to handle modular services using a structured approach. The application includes separate layers for handling requests, middleware for processing, models for representing data, and services for implementing core business logic. It is containerized using Docker to simplify deployment and scalability.

## Features

- **Modular Structure**: Handlers, middleware, models, and services are well-structured to ensure maintainability and scalability.
- **Containerization**: Easily run the project using Docker for a hassle-free development and production environment.
- **Dependency Management**: Utilizes Go modules (`go.mod` and `go.sum`) to manage external libraries and dependencies.

## Prerequisites

Before you begin, make sure you have the following tools installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Docker](https://www.docker.com/products/docker-desktop) (for Docker-based deployment)
- [Git](https://git-scm.com/) (to clone the repository)

## Installation

### 1. Clone the repository

Start by cloning the repository to your local machine:

```bash
git clone https://github.com/yourusername/Xapiens_test.git
cd Xapiens_test
```

### 2. Install Go Dependencies

```bash
go mod tidy
```

### 3. Build the Project

```bash
go build -o xapiens_test
```

### 4. Run the Project

```bash
./xapiens_test
```

## Docker Setup

Build the Docker image:

```bash
docker build -t xapiens_test .
```

Run the container:

```bash
docker run -p 8080:8080 xapiens_test
```
