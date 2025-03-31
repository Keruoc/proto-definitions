# Proto Definitions

This folder contains `.proto` files and their generated code.

## Structure

- `proto/`: Source `.proto` files.
  - `common/`: Reusable definitions (e.g., `Link`, `Node`).
  - `kss/`: Service-specific definitions (e.g., `KnowledgeStorageService`).
- `gen/`: Generated Go code.
- `Makefile`: Automates code generation.

## Generating Code

Run the following command to generate Go code:

```sh
make gen
```
