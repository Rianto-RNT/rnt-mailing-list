## Project Setup

- Microservice with golang
  - JSON API Server
  - gRPC API Server
  - gRPC API Client
  - Protocol Buffers
  - SQLite
  - Gorutines
  - CLI

## Packages
- [go-sqlite3](https://github.com/mattn/go-sqlite3)

## Development Process

### 1) Setup

- Project Microservice

This project requires a `gcc` compiler installed and the `protobuf` code generation tools.

#### Install protobuf compiler

Download the appropriate binary for your operating system from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases).

Install by placing the extracted binary somewhere in your `$PATH`. Additional instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).

#### Install Go protobuf codegen tools

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

#### Generate Go code from .proto files

After creating your `.proto` file, this command will generate the necessary code.

```sh
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  Proto/mail.proto
```

### 2) Creating Database Tables
- Create folder and mdb.go file
- Create email entry function

### 3) Implementing CRUD Operations
- Creating Email function
- Creat a function to get an email
- Creat a function to update an email
- Create a function to delete an email
- Create a function to get batch email

### 4) JSON Data Processing Functions
- Working with JSON API

### 5) JSON Endpoints
- Create CRUD email endpoints function

### 6) Testing JSON API
- Setup server.go
- JSON API Testing with Thunder client (vs code extention)
- Make CRUD emails request

### 7) Protocol Buffers
- New directory and file mail.proto
- Create emails CRUD operation in mail.proto file

### 8) gRPC Data Processing Functions
- Create gRPC API directory and new file grpcapi.go
- Implementing mail.proto in grpcapi
- Create function to specify email entry