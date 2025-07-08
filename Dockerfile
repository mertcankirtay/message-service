FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o message-service .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/message-service .

EXPOSE 8000

ENTRYPOINT ["./message-service"]