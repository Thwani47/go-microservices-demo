FROM golang:1.18-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

CMD ["go", "run", "/app/cmd/web"]

# go run /app/cmd/web