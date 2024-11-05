FROM golang:1.23-alpine AS builder

LABEL app="developer-profile-api" \
      version="1.0" \
      maintainer="your-email@example.com"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
