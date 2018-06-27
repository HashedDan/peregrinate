FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD attraction.json /app/attraction.json
ADD attraction-cli /app/attraction-cli

CMD ["./attraction-cli"]