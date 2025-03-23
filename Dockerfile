FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o chat-server main.go

EXPOSE 8080

CMD ["./chat-server", "server"]
