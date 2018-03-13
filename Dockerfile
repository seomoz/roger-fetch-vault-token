FROM golang:1.10.0-alpine3.7

RUN mkdir /output && \
  apk add --no-cache git && \
  go get github.com/mitchellh/gox && \
  go get -u github.com/kardianos/govendor

COPY vendor/vendor.json /go/src/roger-gk-mesos/vendor/vendor.json
COPY main.go /go/src/roger-gk-mesos/main.go

WORKDIR /go/src/roger-gk-mesos

RUN govendor sync

CMD gox -output='/output/{{.Dir}}_{{.OS}}_{{.Arch}}' -tags='netgo' -ldflags='-w'
