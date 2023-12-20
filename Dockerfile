# Start from a Golang image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Install Node.js and npm
RUN apt-get update && \
    apt-get install -y nodejs npm

# Build the Vue.js frontend
WORKDIR /app/front
RUN npm install
RUN npm run build

WORKDIR /app

RUN go build -o main .

# Final image
FROM golang:latest

WORKDIR /app

# Copy only the necessary files from the builder image
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
