FROM golang:1.10.0-alpine3.7

RUN mkdir /output && \
  apk add --no-cache git && \
  go get github.com/mitchellh/gox && \
  go get -u github.com/kardianos/govendor

COPY vendor/vendor.json /go/src/roger-fetch-vault-token/vendor/vendor.json
COPY main.go /go/src/roger-fetch-vault-token/main.go

WORKDIR /go/src/roger-fetch-vault-token

RUN govendor sync

CMD gox -output='/output/{{.Dir}}_{{.OS}}_{{.Arch}}' -tags='netgo' -ldflags='-w'
