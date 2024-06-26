.PHONY: lint
lint:
	golangci-lint run

.PHONY: local-up
local-up:
	docker-compose -p local -f ./containers/docker-compose.yaml up -d

.PHONY: local-restart
local-restart:
	docker-compose -p local -f ./containers/docker-compose.yaml restart

.PHONY: local-down
local-down:
	docker-compose -p local -f ./containers/docker-compose.yaml down

.PHONY: build
build:
	go build -o ./bin/observability ./cmd/api