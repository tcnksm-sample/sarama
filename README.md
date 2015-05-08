# sarama

This is a sample project to use [sarama](https://github.com/afex/hystrix-go), sarama is golang client library for Apache kafka.


## Install

To get sarama,

```bash
$ go get github.com/Shopify/sarama
```

## Build

To build producer & consumer, Binary is generated in each directory,

```bash
$ make build
```

To build docker image for Apache kafka,

```bash
$ make build-docker
```

## Usage

Run kafka on docker container, 

```bash
$ ./docker-run.sh
```

To execute consumer, run below. Consumer tries to fetch topic until received `os.Interrupt` signal. 

```bash
$ docker exec -it kafka /work/consumer/consumer
```

To execute async-producer, run below. Producer tries to publish messages until received `os.Interrupt` signal. 

```bash
$ docker exec -it kafka /work/async-producer/async-producer
```

## Author

[Taichi Nakashima](https://github.com/tcnksm)


