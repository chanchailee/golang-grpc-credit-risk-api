#!/bin/bash
docker build . -f app.Dockerfile -t generate-fico-records
docker run -v /${PWD}:/go/src/github.com/chanchailee/golang-grpc-credit-risk-api/generate-fico-records/ -it generate-fico-records generate-fico-records $1 $2