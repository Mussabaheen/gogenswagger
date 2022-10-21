test:
	$(info 🧪 Testing...)
	go test ./... -short

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
	@chmod ug+x hooks/*
	@chmod ug+x .git/hooks/*

pre-commit: format lint test build