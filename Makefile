.PHONY: run
run:
	$(GOENV) go run cmd/app/main.go $(RUN_ARGS)

.PHONY: generate
generate:
	$(GOENV) protoc --go_out=. --go-grpc_out=. api/contract.proto

.PHONY: migrate
migrate:
	$(GOENV) go run cmd/migrate/main.go $s(RUN_ARGS)