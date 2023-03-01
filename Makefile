.PHONY: lint
lint:
	golangci-lint run --config .golangci.yaml

