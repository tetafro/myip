.PHONY: dep
dep:
	@ go mod tidy && go mod verify

.PHONY: lint
lint:
	@ golangci-lint run --fix

.PHONY: build
build:
	@ go build -o ./bin/myip .

.PHONY: run
run:
	@ ./bin/myip

.PHONY: docker
docker:
	@ docker build -t ghcr.io/tetafro/myip .
