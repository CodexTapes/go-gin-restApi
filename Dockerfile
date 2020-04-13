# Base Image
FROM golang:1.13-alpine

# Add tools not present in base image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git

# Maintainer Info
LABEL maintainer="Amir <amir.xx8080@gmail.com>"

# Set Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]

