# Use the stable Go 1.23 image
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Only copy go.mod (since you don't have go.sum)
COPY go.mod ./
RUN go mod download

# Copy your source code
COPY . .

# Build the app
RUN go build -o groupie-app .

# Stage 2: Tiny runtime image
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/groupie-app .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080
CMD ["./groupie-app"]