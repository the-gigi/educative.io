FROM golang

RUN go get -u github.com/onsi/ginkgo/ginkgo && \
    go get github.com/onsi/gomega/...

WORKDIR /
RUN git clone https://github.com/the-gigi/educative.io.git
