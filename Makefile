BUILD_DIR = builds
BULLETIN_ARTIFACT = bulletin
BULLETIN_MAIN = main.go
MAC_GOOS=darwin
MAC_GOARCH=amd64

.PHONY: build run clean test

build:
	env GOOS=$(MAC_GOOS) GOARCH=$(MAC_GOARCH) go build -o $(BUILD_DIR)/$(BULLETIN_ARTIFACT) $(BULLETIN_MAIN)

run:
	go run $(BULLETIN_MAIN)

clean:
	rm -rf $(BUILD_DIR)/*

test:
	go test ./test/...
