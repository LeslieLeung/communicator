# Build stage
FROM golang:1.21-alpine as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o communicator .

# Final stage
FROM alpine:3.19.0

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/communicator .

# Expose port 8080 to the outside
EXPOSE 8080

CMD ["/app/communicator", "serve"]