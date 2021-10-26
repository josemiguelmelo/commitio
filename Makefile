.PHONY: go-get
go-get:
	@echo " > Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod vendor

.PHONY: lint
lint:
	@echo " > Linting ..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

.PHONY: test
test:
	@echo " > Testing ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -coverprofile=cover.out ./...

.PHONY: coverage
coverage:
	mkdir -p coverage
	gocover-cobertura < cover.out > coverage/coverage.xml

.PHONY: gofmt
gofmt:
	@echo " > Formatting ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) gofmt -w .
