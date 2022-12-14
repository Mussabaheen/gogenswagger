test:
	$(info ๐งช Testing...)
	go test ./... -short

lint:
	$(info ๐งช Checking the lint)
	golangci-lint run ./...

format:
	$(info ๐๏ธ formatting...)
	@go fmt ./...


build: clean init-hooks
	$(info ๐ฆ Building...)
	go build -o build/ ./...

clean:
	rm -rf ./build/* 

init-hooks:
	@cp -a hooks/. .git/hooks/
	@chmod ug+x hooks/*
	@chmod ug+x .git/hooks/*

pre-commit: format lint test build