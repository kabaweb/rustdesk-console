FROM golang:1.22-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o nexus-server .

FROM alpine:3.19
RUN apk add --no-cache ca-certificates python3 py3-pip sqlite
WORKDIR /app
COPY --from=builder /app/nexus-server .
COPY customize/ ./customize/
RUN mkdir -p /app/data /app/artifacts
EXPOSE 8090
CMD ["./nexus-server"]
