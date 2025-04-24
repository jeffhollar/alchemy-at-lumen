FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
# RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
ENV CGO_ENABLED=0
RUN go build -o act-core-service .

# Use a minimal runtime image
FROM alpine:3.19

WORKDIR /app
COPY --from=build /app/act-core-service .
# Copy SSL certificate files
COPY key.pem server.key ./
# Copy Swagger UI files
COPY swaggerui ./swaggerui/

# Expose the port defined in Makefile
EXPOSE 8501

# Run the application
ENTRYPOINT ["./act-core-service"]
