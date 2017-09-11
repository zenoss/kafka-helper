FROM golang:1.9-alpine

COPY . /go/src/github.com/zenoss/kafka-helper
WORKDIR /go/src/github.com/zenoss/kafka-helper
RUN go build -o /bin/kafka-helper .

ENTRYPOINT ["/bin/kafka-helper"]
