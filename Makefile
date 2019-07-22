# deps ensures fresh go.mod and go.sum.
.PHONY: deps
deps:
	@go mod tidy

.PHONY: build
build:
	@go build -o janus cmd/janus/*.go
