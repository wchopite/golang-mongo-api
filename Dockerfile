FROM golang:latest

LABEL maintainer="Wladimir Chópite <wchopite@sparkdigital.com>"

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 1337

CMD ["./main"]
