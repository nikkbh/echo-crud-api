FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod tidy

