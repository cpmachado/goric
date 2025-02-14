build:
	go build ./cmd/goric

clean:
	rm -f goric

format:
	gofumpt -w .

format-check:
	gofumpt -l .

lint:
	golangci-lint run .

run:
	go run ./cmd/goric


.PHONY: build clean format format-check lint run
