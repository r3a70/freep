FROM golang:1.23.2

WORKDIR /app

COPY . /app

RUN go build .

EXPOSE 8000
