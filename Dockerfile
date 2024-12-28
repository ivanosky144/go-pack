FROM golang:1.22.2 AS builder

WORKDIR /app

# Copy go.mod and go.sum to handle dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the internal directory and all necessary Go files
COPY internal/ ./internal/

# Copy any other files you need (like .yaml files or other configs)
COPY . .

# Print files to ensure they are in the container (for debugging)
RUN ls -R /app

# Build the Go binary from the correct directory
RUN go build -o gopack ./internal/binary_tree

# If you want to build from another directory, change the build command:
# RUN go build -o gopack ./internal/hash_set
# RUN go build -o gopack ./internal/linked_list

# CMD to run the built binary (if needed, or define entry point)
# CMD ["./gopack"]
