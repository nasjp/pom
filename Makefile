build: test
	@echo build start ...
	@go build -o pom cmd/main.go
	@echo created pom!!

test:
	@echo start testing
	@go test ./...
	@echo test was successful!!

clean-docker: clean-docer-comopse-images clean-docker-images clean-docker-containers;

clean-docker-containers:
	-docker rm `docker container ls -qa`

clean-docker-images:
	-docker rmi -f `docker images -qa`

clean-docer-comopse-images:
	-docker-compose down --rmi all

in-docker-container:
	cd cmd/; fresh
