FROM golang:1.12.0-alpine

RUN mkdir /output && \
  apk add --no-cache git && \
  go get github.com/mitchellh/gox && \
  mkdir /app

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY main.go .

CMD gox -output='/output/{{.Dir}}_{{.OS}}_{{.Arch}}' -tags='netgo' -ldflags='-w'
