
# File Reader Web Service

This project is a service that communicates with multiple APIs, retrieves data, and stores it in a database.

Project diagrams can be found at the following link:  
https://drive.google.com/file/d/12TZbFSVDtpbwpe5jJVOaCnM4Of2nH64V/view?usp=sharing  
They can also be viewed in the `code_challenge_diagrams.drawio.pdf` file within this repository.

## Table of Contents

- Description  
- Prerequisites  
- Installation  
- Configuration  
- Usage  
- Project Structure  
- API Endpoints  
- Contributions  

## Description

This service processes files in CSV and JSON Lines formats, interacts with external MercadoLibre APIs to retrieve additional data, and stores the information in a PostgreSQL database. It is containerized using Docker and orchestrated with Docker Compose.

## Prerequisites

- [Docker](https://www.docker.com/)  
- [Docker Compose](https://docs.docker.com/compose/)  
- [Go](https://golang.org/) (for local development)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-repo/project.git
    cd project
    ```

2. Build the Docker containers:

    ```sh
    docker-compose up -d
    ```

## Configuration

Set up the environment variables by editing the following files:

- `config.env`
- `config.go`

### config.env

```env
# File format to process (csv, jsonl)
FILE_FORMAT=jsonl

# Separator for CSV files (e.g., comma, tab, etc.)
FILE_SEPARATOR=,

# File encoding (e.g., utf-8, iso-8859-1)
FILE_ENCODING=utf-8

# Directory where the files to be processed are located
FILE_DIRECTORY=./pending

# Server port
PORT=8080

# Mediator endpoint URL
MEDIATOR_URL=http://mediator:8081/receive-files

BEARER_TOKEN=your_token_here
```

## Usage

To start the application using Docker Compose, run:

```sh
docker-compose up -d
```

Available endpoints:

- `POST /process-files`: Processes the pending files in the configured directory.  
- `POST /receive-files`: Endpoint to receive processed files from the mediator.

## Project Structure

```
.
├── .gitignore
├── docker-compose.yml
├── file_processor/
│   ├── api/
│   │   └── post.go
│   ├── config.env
│   ├── Dockerfile
│   ├── exports
│   ├── file_processor/
│   │   ├── csv_reader.go
│   │   ├── file_reader_template.go
│   │   └── jsonline_reader.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── pending/
│   │   ├── technical_challenge_data_jsonl.jsonl
│   │   └── technical_challenge_data.csv
│   └── utils/
│       └── config.go
├── mediator/
│   ├── apis/
│   │   ├── api.go
│   │   ├── category_API.go
│   │   ├── currency_API.go
│   │   └── user_API.go
│   ├── config.env
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── mediator/
│   │   └── mediator.go
│   ├── processed/
│   ├── test/
│   └── utils/
├── mssql-init/
│   └── init.sql
├── other/
│   └── file_reader/
│       ├── api_factory/
│       │   ├── api_factory.go
│       │   └── repository.go
│       └── utils/
│           └── config.go
├── README.md
└── token_requests
```

## API Endpoints

### file_processor

- **Process Pending Files**

    `POST /process-files`

    Processes the files in the directory specified in `config.env`.

    ```go
    // [file_processor/main.go](file_processor/main.go)
    func main() {
        // ...
        r.POST("/process-files", func(c *gin.Context) {
            // ...
        })
        // ...
    }
    ```

### mediator

- **Receive Processed Files**

    `POST /receive-files`

    Receives the processed files from `file_processor` and stores them in the database.

    ```go
    // [mediator/main.go](mediator/main.go)
    func main() {
        // ...
        r.POST("/receive-files", func(c *gin.Context) {
            // ...
        })
        // ...
    }
    ```

