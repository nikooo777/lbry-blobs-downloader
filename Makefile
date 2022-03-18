BINARY=blobsdownloader

DIR = $(shell cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)
BIN_DIR = ${DIR}/bin
IMPORT_PATH = github.com/nikooo777/lbry-blobs-downloader

VERSION = $(shell git --git-dir=${DIR}/.git describe --dirty --always --long --abbrev=7)
LDFLAGS = -ldflags "-X ${IMPORT_PATH}/meta.Version=${VERSION} -X ${IMPORT_PATH}/meta.Time=$(shell date +%s)"


.PHONY: build clean test lint
.DEFAULT_GOAL: build


build:
	mkdir -p ${BIN_DIR} && CGO_ENABLED=0 go build ${LDFLAGS} -asmflags -trimpath=${DIR} -o ${BIN_DIR}/${BINARY} main.go

clean:
	if [ -f ${BIN_DIR}/${BINARY} ]; then rm ${BIN_DIR}/${BINARY}; fi

test:
	go test ./... -v -cover

lint:
	go install github.com/alecthomas/gometalinter@latest && gometalinter --install && gometalinter ./...