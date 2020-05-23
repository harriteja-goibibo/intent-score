default: clean build

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u github.com/alecthomas/gometalinter
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/dep/cmd/dep
	gometalinter --install --update
	@$(MAKE) dep

.PHONY: dep
dep: ## Run dep ensure and prune
	dep ensure
	dep prune

.PHONY: fmt
fmt: ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint: ## Run all the linters
	gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		./...

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: build
build:  ## Build a version
	mkdir bin
	go build -v -o ./bin/intent-score_server ./cmd/intent-score/main.go

.PHONY: clean
clean: ## Remove temporary files
	rm -rf ./bin
	rm -f coverage.txt
