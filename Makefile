GO ?= go
GO_TAGS ?= go_json
GO_FILES = $(shell find . -name \*.go)
GO_BUILDFLAG ?= -v trimpath$(if $(GO_TAGS), -tags $(GO_TAGS),)
GO_TESTFLAGS ?= -tags "test$(if $)"
GOTESTFLAGS ?= -tags "test$(if $(GO_TAGS),$(comma)$(GO_TAGS),)" -short -race

all: generate build

preparebeta:
	@printf \\e[1m"Create beta env file"\\e[0m\\n
	@$(GO) mod vendor
	@cp -n .env.example .env.beta || true
	@printf \\e[1m"-------------------Finish init beta environment-------------------"\\e[0m\\n

initbeta: preparebeta
	@printf \\e[1m"Start mongodb replicaset instance via docker"\\e[0m\\n
	@chmod +x ./scripts/rs-init.sh
	@docker compose up -d
	@sleep 5
	@docker exec mongo1 /scripts/rs-init.sh
	@printf \\e[1m"Success startup mongodb"\\e[0m\\n

generate: pregenerate
	@printf \\e[1m"Generate"\\e[0m\\n
	@$(GO) generate ./...
	@cd proto && $(GO) generate

pregenerate:
	@printf \\e[1m"Install dependency"\\e[0m\\n
	@$(GO) install github.com/golang/mock/mockgen@v1.6.0

test:
	@printf \\e[1m"Run test"\\e[0m\\n
	@ENV=unittest $(GO) test $(GOTESTFLAGS) ./...

build: .bin/rssi-grpc

go.sum: go.mod
	@printf \\e[1m"go mod tidy"\\e[0m\\n
	@$(GO) mod tidy

.bin/rssi-grpc: go.mod go.sum $(GO_FILES)
	@printf \\e[1m"Build .bin/rssi-grpc"\\e[0m\\n
	@cd cmd/rssi-grpc && $(GO) build -o ../../.bin/rssi-grpc .