FROM golang:alpine as builder

WORKDIR /app 

COPY . .

ENV GIN_MODE=release
RUN cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM busybox

WORKDIR /app

COPY --from=builder /app/src/go-dart-pub /usr/bin/
RUN mkdir /data

ENTRYPOINT ["go-dart-pub"]