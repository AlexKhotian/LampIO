BINARY_NAME=LampIO
default: build

.PHONY: test
test:
	go test -v -race ./...

.PHONY: clean
clean:
	go clean
	rm $(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_NAME) -v ./CLI