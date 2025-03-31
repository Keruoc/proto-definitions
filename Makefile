.PHONY: gen

PROTO_SRC := proto
GEN_OUT := .

gen:
	@echo "Generating Go code from .proto files..."
	protoc \
		-I=$(PROTO_SRC) \
		--go_out=paths=source_relative:$(GEN_OUT) \
		--go-grpc_out=paths=source_relative:$(GEN_OUT) \
		$(shell find $(PROTO_SRC) -name "*.proto")
	@echo "Done generating Go code."