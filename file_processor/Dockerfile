# Stage 1: Build the Go application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy the .env file
COPY config.env ./

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o file_processor_binary main.go

# Stage 2: Run the application in a minimal Alpine image
FROM alpine:latest

RUN apk add --no-cache file

# Copy the .env file and the built binary
COPY --from=builder /app/pending /app/pending
COPY --from=builder /app/config.env /app/config.env
COPY --from=builder /app/file_processor_binary /app/file_processor_binary

# Set execution permissions
RUN chmod +x /app/file_processor_binary

# Set the working directory to /app
WORKDIR /app

EXPOSE 8081

CMD ["/app/file_processor_binary"]