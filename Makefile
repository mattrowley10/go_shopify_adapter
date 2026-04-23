build:
	@go build -o bin/app ./cmd/api

run: build
	@./load_env.sh