# sarama

This is a sample project to use [sarama](https://godoc.org/github.com/Shopify/sarama), sarama is golang client library for Apache kafka. This repositry contains simple producer and consumer. To use this you need to prepare kafka (or you can run it by docker image). 


## Install

To get sarama,

```bash
$ go get github.com/Shopify/sarama
```

## Build

To build producer & consumer, run below. Binary is generated in each directory,

```bash
$ make build
```

## Usage

Ensure kafka is running on where you expect.

To execute consumer, run below. Consumer tries to fetch topic until received `os.Interrupt` signal. 

```bash
$ ./consumer/consumer
```

To execute async-producer, run below. Producer tries to publish messages until received `os.Interrupt` signal. 

```bash
$ ./async-producer/async-producer
```

## Kafka on Docker

You can run kafka on docker. And you can also run producer & consumer on docker.

To build docker image of kafka,

```bash
$ make docker-build
```

To run kafka container, 

```bash
$ ./docker-run.sh
```

To execute consumer,

```bash
$ docker exec -it kafka /work/consumer/consumer
```

To execute async-producer,

```bash
$ docker exec -it kafka /work/async-producer/async-producer
```

To execute http-log-producer,

```bash
$ docker exec -it kafka /work/http-log-producer/http-log-producer
```

And access it (e.g., docker is on boot2docker), 

```bash
$ curl $(boot2docker ip):8080
```

## Author

[Taichi Nakashima](https://github.com/tcnksm)


