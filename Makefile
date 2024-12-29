NAME       			:= fittui
GO_FLAGS   			?=
OUTPUT_BIN 			?= bin/${NAME}
PACKAGE    			:= fittui
GIT_REV    			?= $(shell git rev-parse --short HEAD)
SOURCE_DATE_EPOCH 	?= $(shell date +%s)
ifeq ($(shell uname), Darwin)
	DATE  			?= $(shell TZ=UTC date -j -f "%s" ${SOURCE_DATE_EPOCH} +"%Y-%m-%dT%H:%M:%SZ")
else
	DATE		    ?= $(shell date -u -d @${SOURCE_DATE_EPOCH} +"%Y-%m-%dT%H:%M:%SZ")
endif
VERSION    			?= v0.0.1

default: help

test:   ## Run all tests
	@go clean --testcache && go test ./...

cover:  ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

build:  ## Builds the iaws-fargate-driver executable
	@go build ${GO_FLAGS} \
	-ldflags "-w -s -X ${PACKAGE}/cmd.version=${VERSION} -X ${PACKAGE}/cmd.commit=${GIT_REV} -X ${PACKAGE}/cmd.buildDate=${DATE}" \
	-a -o ${OUTPUT_BIN} main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'