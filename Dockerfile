FROM go:1.23.2

WORKDIR /app

copy . /app

RUN go build .

EXPOSE 8000
