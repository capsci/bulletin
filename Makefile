BUILD_DIR = builds
BULLETIN_ARTIFACT = bulletin
BULLETIN_MAIN = main.go

build:
	go build -o $(BUILD_DIR)/$(BULLETIN_ARTIFACT) $(BULLETIN_MAIN)

run:
	go run $(BULLETIN_MAIN)

clean:
	rm -rf $(BUILD_DIR)/*
