
test: test-unit test-coverage

test-unit:
	@echo "=== Running unit test ==="
	go test -v -race -short ./...

test-coverage:
	@echo "=== Running coverage test ==="
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	rm -f coverage.out
	@echo "Coverage report generated: coverage.html"
