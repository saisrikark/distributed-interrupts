.PHONY: all clean

BIN_DIR = bin

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

all: noderunner

noderunner: cmd/noderunner/main.go
	go build -o $(BIN_DIR)/noderunner ./cmd/noderunner

clean:
	rm -rf $(BIN_DIR)
