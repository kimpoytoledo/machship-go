

# Github API - Golang

## Table of Contents

- [Overview](#overview)
- [Architecture Overview](#architecture-overview)
- [Project Structure](#project-structure)
- [Key Dependencies](#key-dependencies)
- [Usage Guide](#usage-guide)
  - [Prerequisites](#prerequisites)
  - [Instructions](#instructions)
- [Advanced Features](#advanced-features)

## Overview

This project serves as an API written in Go that fetches Github user data. It uses a Redis cache to store and retrieve Github user data swiftly for subsequent requests, ensuring enhanced performance. The Hexagonal (or Ports and Adapters) design pattern underpins the architecture, emphasizing a clear separation of concerns and adaptability.

## Architecture Overview

The application adopts the Hexagonal Architecture, which segregates it into distinct sections:

- **Core Domain:** Contains essential business logic and entities.
- **Ports:** Interfaces that establish expected interactions and behaviors.
- **Adapters:** Implementations of the ports that facilitate interaction with external systems or services.

## Project Structure

```
machship-go/
│
├── adapter/                  # Adapters for external services
│   ├── redis/                # Redis adapter
│   └── githubapi/            # Github API adapter
│
├── api/                      # API-related code (endpoints, handlers)
│
├── cmd/                      # Main application entry point
│
├── core/                     # Core domain logic
│   ├── entity/               # Entity definitions
│   └── usecase/              # Core business logic
│
├── util/                     # Utilities (logging, error handling)
│   ├── logger/
│   └── errorhandler/
│
├── docker-compose.yml        # Docker Compose configuration
└── Dockerfile                # Dockerfile for the Go application
```

## Key Dependencies

- **Gin:** A lightweight web framework utilized for crafting the API.
- **Go Redis:** A Redis client tailored for Go, facilitating interaction with the Redis cache.
- **Github API:** An external service supplying comprehensive Github User data.

## Usage Guide

### Prerequisites:

Before diving in, ensure Docker and Docker Compose are installed:

- [Docker installation guide](https://docs.docker.com/get-docker/)
- [Docker Compose installation guide](https://docs.docker.com/compose/install/)

### Instructions:

1. **Clone the Repository:**

   ```bash
   git clone <repository-url>
   cd machship-go
   ```

2. **Activate the Services:**

   Initiate the build and startup process using Docker Compose:

   ```bash
   docker-compose up --build
   ```

3. **Engage with the API:**

   Once the services are active, the API can be accessed using the following `curl` command:

   ```bash
   curl --location 'http://localhost:8080/github' \
        --header 'Content-Type: application/json' \
        --data '{
           "usernames" : ["random1111","octocat","kimpoytoledo"]
        }'
   ```

   In the provided command, the `usernames` field contains a list of Github usernames you wish to fetch data for. Feel free to replace the example usernames (`"random1111","octocat","kimpoytoledo"`) with your desired Github usernames. The resulting response will provide the Github User data, indicating the data source as either "Redis Cache" or "GithubAPI".


4. **Deactivate the Services:**

   To gracefully halt the services, use `CTRL+C` in the terminal. For a complete shutdown and removal of containers, execute:

   ```bash
   docker-compose down
   ```

## Advanced Features

- **Logging:** The integrated logging utilities provide insights into significant events and potential errors, aiding in troubleshooting and understanding the application's behavior.
- **Error Handling:** A dedicated utility for error handling ensures consistency and extensibility in managing and reporting errors.


