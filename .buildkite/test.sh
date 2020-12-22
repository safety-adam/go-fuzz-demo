#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

docker run --network host \
    -v $(pwd):/src \
    -w /src \
    303305260587.dkr.ecr.us-east-1.amazonaws.com/safetyculture/golang:1.15 \
    /bin/bash \
    -c "go build && \
    go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build && \
    go-fuzz-build ./strfuncs && \
    timeout 360 go-fuzz -bin=./strfuncs-fuzz.zip -func=FuzzConvertToBool"
