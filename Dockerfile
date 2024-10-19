# Stage 1: Build the Go binaries using the official Golang image
FROM golang:1.22 AS builder

WORKDIR /app

# Environment variables
ENV API_PORT=3000
ENV TOKEN_SECRET=asdknasjdbakjbdiuawbeiybajkdnkasndmhkbfihagwiura
ENV TOKEN_TIME_ACCESS=150
ENV TOKEN_TIME_REFRESH=168
ENV TOKEN_REMEMBER_REFRESH=8760
ENV DB_NAME=defaultdb
ENV DB_HOST=didlydoodash-db-didlydoodash.k.aivencloud.com
ENV DB_PORT=14575
ENV DB_USER=avnadmin
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_SSL=require
ENV DB_TIMEZONE=Europe/Helsinki
ENV MODE=production

# Copy go.mod and go.sum to install dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source files and certificate
COPY . .

# Build the Go binaries for the migrate and api commands
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./didlydoodash-migrate ./src/cmd/migrate && \
    CGO_ENABLED=0 GOOS=linux go build -v -o ./didlydoodash-api ./src/cmd/api

# Stage 2: Create a lightweight image using Alpine Linux
FROM alpine:latest

WORKDIR /api

# Environment variables
ENV API_PORT=3000
ENV TOKEN_SECRET=asdknasjdbakjbdiuawbeiybajkdnkasndmhkbfihagwiura
ENV TOKEN_TIME_ACCESS=150
ENV TOKEN_TIME_REFRESH=168
ENV TOKEN_REMEMBER_REFRESH=8760
ENV DB_NAME=defaultdb
ENV DB_HOST=didlydoodash-db-didlydoodash.k.aivencloud.com
ENV DB_PORT=14575
ENV DB_USER=avnadmin
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_SSL=require
ENV DB_TIMEZONE=Europe/Helsinki
ENV MODE=production

# Copy the built binaries and the startup script from the builder stage
COPY --from=builder /app/didlydoodash-api /app/didlydoodash-migrate /app/start.sh /app/ca.pem ./

# Ensure the script has unix line endings
RUN sed -i 's/\r$//' start.sh

# Give execution permissions to the startup script
RUN chmod +x start.sh

# Expose the port the API will run on
EXPOSE 3000

# Set the startup script as the container's entry point
ENTRYPOINT ["sh","./start.sh"]