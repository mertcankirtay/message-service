FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/message-service .

EXPOSE 8000

ENTRYPOINT ["./message-service"]