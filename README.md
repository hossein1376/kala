# KALA

Kala (meaning goods) is an online shop backend written in Go.  
It aims to be fast, simple and extensible.

## Ent

### Generate Ent

```shell
go generate ./internal/ent
```

### Generate new schema file

```shell
go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema <Name>
```