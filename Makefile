fetch=true

.PHONY: tidy
## tidy: download lib dependencies
tidy:
	go mod tidy

.PHONY: build
## build: build project files
build: tidy
	go build .

.PHONY: run
## run: run project
run: tidy
	go run main.go -tags=$(tags) -fetch=$(fetch)

.PHONY: test
## run: run project
test2: tidy
	go test ./services -v -coverprofile .cov
	go tool cover -func=.cov | grep "total:"



