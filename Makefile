BINARY_NAME=joker
SRC=main.go
 
all: build test
 
build:
	go build -o ${BINARY_NAME} $(SRC)
 
test:
	go test
 
run: build
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}
