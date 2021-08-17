FROM golang:1.16.7-alpine as builder

WORKDIR /go/build
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o app

FROM alpine:latest
WORKDIR /home
# Copy the binary file from the first image
COPY --from=builder /go/build/app .
# Copy the env file
COPY .env /home
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./app"]
