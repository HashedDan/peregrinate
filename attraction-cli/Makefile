build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t attraction-cli .

run:
	docker run -e MICRO_REGISTRY=mdns attraction-cli