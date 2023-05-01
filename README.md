# Kafka Go Experiment

A small project to play around with Kafka and Go, consisting of a producer module and a consumer module. Everything is dockerized and can be run with `docker-compose`.

The events being sent into Kafka are fake "purchase" events, they include a unique uuid, buyer's name, item and price. They are serialized to json before being sent.

## Showcase

Below are the logs of two docker containers, showing the producer sending events and the consumer receiving them.

**Top** - Consumer, **Bottom** - Producer

https://user-images.githubusercontent.com/31768805/235538617-2ce80cc8-cf59-4f44-976c-3cc4ef5acc32.mp4

### Modules

There are three modules in this project:

1. `producer` - sends events to Kafka
2. `consumer` - receives events from Kafka
3. `util` - contains shared utility code

Each one has a seperate `Dockerfile`. The expected way to run them is with `docker-compose`, as it contains the build context.

## Development

1. Run `docker-compose up`

2. Rerun `consumer` or `producer` as appropriate, using `docker-compose run --build <consumer|producer>`
