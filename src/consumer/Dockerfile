# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY src/consumer/ consumer/

COPY src/util/ util/

COPY kafka-experiment.properties consumer/

WORKDIR /app/consumer/

RUN go build -o consumer .

CMD ["./consumer", "kafka-experiment.properties"]