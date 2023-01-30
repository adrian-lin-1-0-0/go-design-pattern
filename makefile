.PHONY: test
test:
	go test ./...
.PHONY: lint
lint:
	golangci-lint run ./...
.PHONY: docker-lint
docker-lint:
	docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v