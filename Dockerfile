FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080 8081

VOLUME ["./config", "./config"]

CMD ["./main"]

