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

# Build the Go application with embedded Vue.js frontend
RUN go build -o main .

# Build the Vue.js frontend
WORKDIR /app/front
RUN npm install
RUN npm run build

# Final image
FROM golang:latest

WORKDIR /app

# Copy only the necessary files from the builder image
COPY --from=builder /app/main .
COPY --from=builder /app/front/dist ./front/dist

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
