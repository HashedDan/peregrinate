FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/HashedDan/peregrinate/attraction-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY attraction.json /app/attraction.json
COPY --from=builder /go/src/github.com/HashedDan/peregrinate/attraction-cli/attraction-cli .

ENTRYPOINT ["./attraction-cli", "attraction.json"]
CMD ["./attraction-cli"]