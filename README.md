# Nice Wildcat

## Requirements

- **[Go](https://go.dev):** recommended to use v1.21+ though this should work with v1.20 as well.
- **[Taskfile](https://taskfile.dev):** recommended to use v3+.

## Usage

If using **Taskfile** you can use the single command `task dev`. Else the following commands will work from root.

```bash
# will run the generate script inside `ui/`
go generate ./...
# will run the main script inside `cmd/main`
go run ./...
```