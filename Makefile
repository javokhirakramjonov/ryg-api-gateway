SWAGGER := $(shell which swag)
SWAGGER_OUT_DIR := docs
SWAGGER_GEN_SCRIPT := $(SWAGGER) init -g ./cmd/main.go -o $(SWAGGER_OUT_DIR) --parseDependency --parseInternal --parseDepth 1

swag-gen:
	$(SWAGGER_GEN_SCRIPT)
