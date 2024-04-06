FROM golang:latest

WORKDIR /go/app

COPY . .

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null", "&&", "go", "run", "cmd/consumer/main.go"]