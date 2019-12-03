FROM golang:latest

ENV GOBIN /go/bin
ENV GOPATH /go

# build directories
WORKDIR /go/src/
RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/chanchailee
RUN mkdir /go/src/github.com/chanchailee/golang-grpc-credit-risk-api

WORKDIR /go/src/github.com/chanchailee/golang-grpc-credit-risk-api
ADD . .

#Install additional package
#RUN go get -u github.com/joho/godotenv
#RUN go get -u github.com/stretchr/testify/assert

WORKDIR /go/src/github.com/chanchailee/golang-grpc-credit-risk-api/generate-fico-records
#Go build
RUN go install .