BIN_DIR ?= ./bin
COMMIT ?= $(shell git log -1 --pretty=format:"%h")

build:
	CGO_ENABLED=0 go build -o  "$(BIN_DIR)/hello" .

docker-build:
	docker build --tag edp/hello:$(COMMIT) .

test: 
	go test -cover -failfast
