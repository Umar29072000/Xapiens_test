# Xapiens_test

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
go mod download
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
