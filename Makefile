run:
	make swag-gen
	go run cmd/main.go

SWAGGER := $(shell which swag)
SWAGGER_OUT_DIR := docs
SWAGGER_GEN_SCRIPT := $(SWAGGER) init -g ./api/router.go -o $(SWAGGER_OUT_DIR) --parseDependency --parseInternal --parseDepth 1

swag-gen:
	$(SWAGGER_GEN_SCRIPT)

gen-proto:
	./scripts/genProto.sh .
