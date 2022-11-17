FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /golang-api-test

EXPOSE 8080

CMD [ "/golang-api-test" ]