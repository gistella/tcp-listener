FROM golang:alpine AS build

WORKDIR /tmp/build

COPY . .
RUN go build -v -o tcp-listener ./...

FROM build

COPY --from=build /tmp/build/tcp-listener /usr/local/bin/tcp-listener

CMD ["tcp-listener"]
