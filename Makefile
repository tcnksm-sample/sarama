build:
	cd sync-producer; GOOS=linux GOARCH=amd64 go build
	cd async-producer; GOOS=linux GOARCH=amd64 go build
	cd http-log-producer; GOOS=linux GOARCH=amd64 go build
	cd consumer; GOOS=linux GOARCH=amd64 go build

docker-build:
	cd dockerfile-kafka; docker build -t tcnksm/kafka .
