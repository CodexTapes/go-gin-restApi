# Base Image
FROM golang:1.13.0-alpine3.10 as builder

# Maintainer Info
LABEL maintainer="Amir <amir.xx8080@gmail.com>"

# Set environmet variables for image
# ENV GO111MODULE=auto \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# Add tools not present in base image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git

# Set Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Copy the .env file.
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./main"]