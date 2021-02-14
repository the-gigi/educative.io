FROM golang

WORKDIR /

ENV GO111MODULE=on

RUN apt-get update && apt-get install -y vim \
                   && apt-get install -y tree

RUN go get -u github.com/onsi/ginkgo/ginkgo && \
    go get -u github.com/onsi/gomega/... && \
    go get -u github.com/blang/semver && \
    go get -u github.com/mitchellh/go-homedir && \
    go get -u github.com/spf13/cobra && \
    go get -u github.com/spf13/viper && \
    go get -u github.com/google/go-github/v33 && \
    go get -u github.com/rhysd/go-github-selfupdate/...

RUN git clone https://github.com/the-gigi/educative.io.git


