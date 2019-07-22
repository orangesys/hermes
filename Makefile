# deps ensures fresh go.mod and go.sum.
.PHONY: deps
deps:
	@go mod tidy
