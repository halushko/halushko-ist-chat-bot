FROM golang:1.25 AS builder
WORKDIR /app
RUN go mod init github.com/halushko/halushko-ist-chat-bot
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /app/halushko-ist-chat-bot

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/halushko-ist-chat-bot .
RUN chmod +x /root/halushko-ist-chat-bot
CMD ["./halushko-ist-chat-bot"]
