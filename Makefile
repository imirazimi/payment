
ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

build:
	go build -o $(ROOT)/bin/payment $(ROOT)/cmd

run:
	go run $(ROOT)/cmd