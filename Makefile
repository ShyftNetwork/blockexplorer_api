PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... exclude=gosec --vendor

run:
	@echo "  >  \033[32mStarting server...\033[0m "
	go run *.go

build:
	@echo "  >  \033[32mBuilding binary...\033[0m "
	go build -o shyft_api

start:
	@echo "  >  \033[32mStarting server...\033[0m "
	./shyft_api

install:
	@echo "  >  \033[32mInstalling dependencies...\033[0m "
	go mod vendor
