build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/hasheddan/peregrinate/attraction-service \
	  proto/attraction/attraction.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t attraction-service .
run:
	docker run -p 50051:50051 attraction-service