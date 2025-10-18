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

.PHONY: deploy
deploy:
	@ ansible-playbook \
	--private-key ~/.ssh/id_ed25519 \
	--inventory '${SSH_SERVER},' \
	--user ${SSH_USER} \
	./playbook.yml
