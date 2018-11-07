default: format
	@go run main.go

test: format
	@go test

format:
	@gofmt -w main.go

install:
	go install

clean:
	go clean

.PHONY: clean
