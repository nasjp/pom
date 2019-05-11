# CONST
GORUN=go run
GOBUILD=go build
GOCLEAN=go clean ./...
GOTEST=go test ./...
DOCKER=docker exec -it
BINARY_NAME=pom
DOCKER_SERVICE_NAME=pom_app_1
MAIN=cmd/main.go

# VAR
var=

all: build

build: test clean
	${GOBUILD} -o ${BINARY_NAME} ${MAIN}

test:
	${GOTEST}

clean:
	${GOCLEAN}

docker:
	${DOCKER} ${DOCKER_SERVICE_NAME} ${GORUN} ${MAIN} ${var}

clean-docker: clean-docer-comopse-images clean-docker-containers clean-docker-images;

clean-docker-containers:
	@echo ============================ start docker rm ===========================
	-@docker rm `docker container ls -qa`

clean-docker-images:
	@echo =========================== start docker rmi ===========================
	-@docker rmi -f `docker images -qa`

clean-docer-comopse-images:
	@echo ======================= start docker-compose down ======================
	-@docker-compose down --rmi all


# in docker container

in-docker-container:
	cd cmd/; fresh
