# STAGE 1: Build the binary
FROM golang:1.23-alpine AS builder 

WORKDIR /app

# Only copy go.mod since go.sum doesn't exist
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app
RUN go build -o groupie-app .

# STAGE 2: Create the small runtime image
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/groupie-app .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080
CMD ["./groupie-app"]