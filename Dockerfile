# Build stage
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /main

# Run stage
FROM scratch
COPY --from=builder /main /main
EXPOSE 8080
CMD ["/main"]