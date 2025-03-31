# Proto Definitions

This folder contains `.proto` files and their generated code. You can use this as a Go package.

## Table of Contents

- [Structure](#structure)
- [Generating Code](#generating-code)
- [Setup](#setup)
- [Installing as a Go Package](#installing-as-a-go-package)

---

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

## Setup

Perform this if you wish to generate the proto files on your own, else you may skip this step.

### Install Protocol Buffer Compiler

Install Protocol buffer compiler

```bash
brew install protobuf
```

### Install Go Plugins for Protobuf

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Installing as a Go Package

To use this as a Go package in your project, run:

```bash
go get github.com/keruoc/proto-definitions@<version-tag>
```

Replace `<version-tag>` with the desired version (e.g., `v1.0.0`).

### Example of using this as a package as the server (NOT client)

```go
package grpc

import (
	"context"

	commonv1 "github.com/keruoc/proto-definitions/common/v1"
	kssv1 "github.com/keruoc/proto-definitions/services/kss/v1"
)

type myGRPCServer struct {
	kssv1.UnimplementedKnowledgeStorageServiceServer // forward compatibility
}

// Constructor
func NewMyGRPCServer() kssv1.KnowledgeStorageServiceServer {
	return &myGRPCServer{}
}

func (s *myGRPCServer) GetSuggestedNodesAndLinks(
	ctx context.Context,
	req *kssv1.KSSRequest,
) (*kssv1.KSSResponse, error) {
	// TODO: use implement real logic
	return &kssv1.KSSResponse{
		SimilarNodes: []*kssv1.SimilarNode{
			{
				Id:              "demo-node",
				Title:           "Demo Node",
				SimilarityScore: 0.99,
			},
		},
		CurrentLink: []*commonv1.Link{},
	}, nil
}
```

---
