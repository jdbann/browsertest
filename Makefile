## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | sort -d | column -t -s ':' |  sed -e 's/^/ /'

## ci: run tests
.PHONY: ci
ci: test

## test: run tests
.PHONY: test
test:
	go test ./...

## test/cover: collect coverage whilst running tests and present results
.PHONY: test/cover
test/cover:
	@mkdir tmp || true # go test cannot create tmp folder
	go test -coverprofile=tmp/coverage.out ./...
	go tool cover -html=tmp/coverage.out
