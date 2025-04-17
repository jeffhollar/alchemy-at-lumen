# Copyright 2025 Lunmen Technologies, Inc.
# All rights reserved.
#

.PHONY: build package run debug clean

BUILD=1.0.0
PORT=8501
export GO111MODULE=on

build:
	docker build -t comapi-${BUILD} .

package: build
	docker save comapi-${BUILD} | gzip > comapi-${BUILD}.tar.gz
run:
	docker run -d -p ${PORT}:${PORT} --name communication-service comapi-${BUILD}
debug:
	docker run -it -p ${PORT}:${PORT} --name communication-service comapi-${BUILD} /bin/bash

clean:
	docker rm -f communication-service
	docker rmi -f comapi-${BUILD}:latest
