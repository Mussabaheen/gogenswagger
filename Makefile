test:
	$(info 🧪 Testing...)
	go test ./... -short

run:
	$(info 🧪 Running the gogenapi)
	go run cmd/go-test-api/main.go

lint:
	$(info 🧪 Checking the lint)
	golangci-lint run ./...

format:
	$(info 🖊️ formatting...)
	@go fmt ./...


build: clean init-hooks
	$(info 📦 Building...)
	go build -o build/ ./...

clean:
	rm -rf ./build/* 

init-hooks:
	@cp -a hooks/. .git/hooks/

pre-commit: format lint test build