FROM golang:1.19.3-alpine

WORKDIR /app
COPY . .

RUN go build -o main main.go
EXPOSE 8080

CMD ["/app/main"]
