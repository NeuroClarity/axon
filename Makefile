.DEFAULT_GOAL := help
APP?=axon
COMMIT_SHA?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
REGISTRY_REPO?=471943556279.dkr.ecr.us-west-1.amazonaws.com
LOG_FILE_NAME?=./logs/$(shell date +%Y%m%d_%H%M%S)_$(APP).log

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

clean:
		@echo "$(OK_COLOR)==> Cleaning working directory... $(NO_COLOR)"
		@go clean
		rm -f $(APP)

build: clean
		@echo "$(OK_COLOR)==> Building binary... $(NO_COLOR)"
		go build -o $(APP) cmd/$(APP)/main.go

run:
		go run cmd/$(APP)/main.go

run-with-logs:
		@go run -race cmd/$(APP)/main.go 2>&1 | tee -a "$(LOG_FILE_NAME)"

docker-build: build
		docker build -t ${APP} .
		docker tag ${APP} ${APP}:${COMMIT_SHA}
		docker tag ${APP}:${COMMIT_SHA} ${REGISTRY_REPO}/${APP}:${COMMIT_SHA}

docker-push:
		docker push ${REGISTRY_REPO}/${APP}:${COMMIT_SHA}

ecr-login:
		@echo "Logging in to ECR..."
		@$$(aws ecr get-login --no-include-email)

docker-build-and-push: docker-build docker-push
