build: test
	@echo build start ...
	@go build -o pom cmd/main.go
	@echo created pom!!

test:
	@echo start testing
	@go test ./...
	@echo test was successful!!
