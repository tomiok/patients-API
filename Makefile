# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=patients-api
VERSION=0.1.0

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Build step, generates the binary.
build:
	go build ${LDFLAGS} -o ${BINARY} server/main.go

# Web is a mask to run the web interface, in our case the main function will start the http server.
web:
	@clear
	@go run server/main.go

# Run go formatter
fmt:
	@gofmt -w .

# Download the go lint. Not running anything.
lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

# Run the lint across all the project. See more options https://raw.githubusercontent.com/golangci/golangci-lint .
lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

# Run the test for all the directories.
test:
	@clear
	@go test -v ./...

###################
# Docker commands #
###################
up:
	docker-compose up

down:
	docker-compose down --remove-orphans

clean:
	sudo rm -rf db/data
