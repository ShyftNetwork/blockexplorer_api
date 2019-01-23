PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	@echo "  >  \033[32mGetting Linter ready...\033[0m "
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	@echo "  >  \033[32mStarting linter.../033[0m "
	gometalinter -e $(go env GOROOT) exclude=gosec --vendor

run:
	@echo "  >  \033[32mStarting server...\033[0m "
	go run *.go

build:
	@echo "  >  \033[32mBuilding binary...\033[0m "
	go build -o blockx_api

start:
	@echo "  >  \033[32mStarting server...\033[0m "
	./blockx_api

install:
	@echo "  >  \033[32mInstalling dependencies...\033[0m "
	go mod vendor

docker:
	@echo "  >  \033[32mStarting...\033[0m "
	make build
	@echo "  >  \033[32mBuilding docker image from Dockerfile...\033[0m "
	docker build -t shyft_api .
	@echo "  >  \033[32mRunning application through docker...\033[0m "
	docker run -p 8080:8080 shyft_api