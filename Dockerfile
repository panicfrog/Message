FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN export GOPROXY=https://goproxy.cn && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main .

COPY --from=builder /app/config ./config

RUN chmod +x /app/main

EXPOSE 8080 8081

ENTRYPOINT ["/app/main"]

