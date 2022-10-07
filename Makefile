test:
	$(info 🧪 Running all the tests)
	go test ./... -p 1

run:
	$(info 🧪 Running the gogenapi)
	go run cmd/go-test-api/main.go
