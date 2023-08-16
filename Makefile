build:
	@go build -o bin/azureql

run: build
	@./bin/azureql