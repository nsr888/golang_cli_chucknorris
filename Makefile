BINARY_NAME=joker
SRC=main.go
 
all: build test
.PHONY: all
 
build:
	go build -o ${BINARY_NAME} $(SRC)
.PHONY: build
 
test:
	go test
.PHONY: test
 
run: build
	./${BINARY_NAME}
.PHONY: run
 
clean:
	go clean
	rm ${BINARY_NAME}
.PHONY: clean

lint:
	gofumpt -w -s .
	golangci-lint run --fix
.PHONY: lint
