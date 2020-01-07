FROM golang:1.13.5
EXPOSE 8080

RUN apt-get -y -qq update && \
    apt-get -y -qq install build-essential git-core binutils bison gcc make < /dev/null > /dev/null && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

ENV GO_VERSION=1.13.5
ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ENV APPLICATION_PACKAGE="src/api"

ENV IGNORE_GO_GET="false"

WORKDIR /go/src/app
COPY src/api/ .

RUN echo "Running go build" && \
  go build && \
  echo "Running go test" && \
  go test

RUN echo "Running app:"

CMD [ "go", "run" ]
