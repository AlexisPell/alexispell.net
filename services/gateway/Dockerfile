# Start from the official Go image
FROM --platform=linux/amd64 golang:1.22-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# ENV GOOS=linux
# ENV GOARCH=amd64

# Build the Go app
RUN go build -o /go-app ./cmd/gateway/main.go

# Start a new stage from scratch
FROM --platform=linux/amd64 alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=build /go-app .

ENV PORT 80
ENV ENV prod

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./go-app"]