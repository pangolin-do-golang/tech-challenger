# 1 - Build Stage
FROM golang:1.22.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web cmd/rest/main.go

# 2 - Runtime Stage
FROM scratch

# Copy only the binary from the build stage to the final image
COPY --from=builder /app/web /

#Expose ports
EXPOSE 8080

#Run the web usecases on container startup
CMD ["./web"]