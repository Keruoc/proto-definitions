# Proto Definitions

This folder contains `.proto` files and their generated code.

## Structure

- `proto/`: Source `.proto` files.
  - `common/`: Reusable definitions (e.g., `Link`, `Node`).
  - `kss/`: Service-specific definitions (e.g., `KnowledgeStorageService`).
- `Makefile`: Automates code generation.
- Other files could be generated

## Generating Code

Run the following command to generate Go code:

```sh
make gen
```

## Instructions to start

Install Protocol buffer compiler

```bash
brew install protobuf
```

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```
