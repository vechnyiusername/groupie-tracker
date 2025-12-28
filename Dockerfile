# STAGE 1: Build the binary
FROM golang:1.25-alpine AS builder

# Install git (needed for some go modules)
RUN apk add --no-cache git

WORKDIR /app

# Copy dependency files first (for faster caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app (output name: groupie-app)
RUN go build -o groupie-app .

# STAGE 2: Create the small runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# IMPORTANT: Copy the binary AND your web folders from the builder stage
COPY --from=builder /app/groupie-app .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Expose the port your server runs on (usually 8080 or 4000)
EXPOSE 8080

# Run the app
CMD ["./groupie-app"]