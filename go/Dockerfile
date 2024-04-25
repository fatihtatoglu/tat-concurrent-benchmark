# Alpine Linux tabanlı Go image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
