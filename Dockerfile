FROM golang:1.13.5
EXPOSE 8080

RUN apt-get -y -qq update && \
    apt-get -y -qq install build-essential git-core binutils bison gcc make < /dev/null > /dev/null && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

ENV GO_VERSION=1.13.5
ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

ENV IGNORE_GO_GET="false"

WORKDIR /go/src/github.com/arpb2/C-3PO
COPY . ./

RUN echo "Checking go fmt formatting" && \
  gofmt -l api/ build/ cmd/ hack/ pkg/  | if [ $(grep -c -o -E ".*") -gt 0 ]; then exit 1; fi

RUN echo "Running go build" && \
  go build ./...

RUN echo "Running go test" && \
  go test ./...

CMD [ "go", "run", "cmd/main.go" ]
