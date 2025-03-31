.PHONY: gen

gen:
	@mkdir -p gen
	@echo "Generating Go code from proto files..."
	protoc \
		-I=proto \
		--go_out=gen \
		--go-grpc_out=gen \
		$(shell find proto -name "*.proto")
	@echo "Done generating Go code from proto files!"

clean:
	@echo "Cleaning up generated files..."
	@rm -rf gen
	@echo "Done cleaning up generated files!"