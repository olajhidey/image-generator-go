# Base image
FROM golang:1.19

# Working directory for your code
WORKDIR /app

# Copy our work file to the /app directory
COPY . .

# Download the necessary modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build

EXPOSE 8080

# Run the app
CMD ["/build"]
