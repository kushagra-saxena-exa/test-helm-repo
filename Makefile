BINARY = bin/sample-http-server
GOOS=linux #$(shell go env GOOS)
GOARCH=arm64 #$(shell go env GOARCH)

.PHONY: clean
clean:
	rm -rfv $(BINARY)

# Build operator binary
.PHONY: build
build: clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY) main.go

.PHONY: test
test:
	go test ./...