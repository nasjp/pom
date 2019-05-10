build: test
	@echo build start ...
	@go build -o pom cmd/main.go
	@echo created pom!!

test:
	@echo start testing
	@go test ./...
	@echo test was successful!!

docker-clean: d-c-clean-image docker-clean-container docker-clean-image;

docker-clean-container:
	-docker rm `docker container ls -qa`

docker-clean-image:
	-docker rmi -f `docker images -qa`

d-c-clean-image:
	-docker-compose down --rmi all

in-docker-container:
	cd cmd/; fresh
