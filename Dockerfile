FROM golang:alpine

WORKDIR /usr/src/tcp-listener

COPY . .
RUN go build -v -o /usr/local/bin/tcp-listener ./...

CMD ["tcp-listener"]
