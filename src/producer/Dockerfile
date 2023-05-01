# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY src/producer/ producer/

COPY src/util/ util/

COPY kafka-experiment.properties producer/

WORKDIR /app/producer/

RUN go build -o producer .

CMD ["./producer", "kafka-experiment.properties"]