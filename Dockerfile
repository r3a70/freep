FROM golang:1.23.2

WORKDIR /app

COPY . /app

RUN go build -o /build/fsp .

EXPOSE 8000
