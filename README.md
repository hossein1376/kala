# KALA

Kala (meaning goods) is an online shop backend written in Go.  
It aims to be fast, simple and extensible. As well as a showcase of my personal project structure.

## Project Structure

The general structure follows the guidelines I picked up
from [Let's Go Further](https://lets-go-further.alexedwards.net/),
but over time I found some shortcoming in that approach which left me to explore for my desired structure.  
The following is the result of my experience of developing multiple Go projects, what I find more logical and extensible
approach,
compile-time checks (I'm looking at you, import cycle), my constant experiment with different designs and on top of all,
my personal taste.

While I find this overall structure to make sense the most, feel free to adapt it partially or completely for your
projects!

### In a look

```
.
├── cmd
|   ├── api
|   |   ├── run.go
|   ├── handlers
|   |   ├── routes.go
|   |   ├── ...
|   ├── state.go
|
├── internal
|   ├── data
|   |   ├── Models.go
|   |   ├── Mocks.go
|   |   ├── ...
|   ├── structure
|   ├── Errors
|
├── pkg
|   ├── json
|   ├── ...
|
├── main.go
├── go.mod
```

### cmd/api

Module `api` is tasked with retrieving the configurations, validation them, opening database connection and finally
starting the server.  
It also handles the graceful shutdown as well.

Configurations, logger and an instance of `Model` will be passed to `handlers`.

### cmd/handlers

`handlers` module feature the struct `handler` which all handlers are a receiver function to it. It has a single
exported receiver function named `Router` witch will be called inside the `cmd/api/run.go` to instantiate the router.

`handlers` module will embed multiple structs such as `Json`, `Errors` so they can be used directly inside the handlers.

### cmd/state.go

This file consist of multiple struct to define the configurations and the "Application" data that will be passed down
to `handlers`
from `cmd`.

The main reason to have this file here, instead of somewhere like `internal/structure`, is to avoid import cycle.

### internal/data

Database queries and transactions will take place here. Each table features a struct with db methods attached to it,
they all will be initialized with the `Models` struct by a database connection.  
With the use of interfaces, it's easy to write mocks for test cases.

### internal/structure

This module consists of structs to define the request, response and data structure, hence the name.

### internal/Errors

HTTP error responses reside here, as well as custom defined errors that mostly will be returned
from `internal/data` to be switch-case-ed in `cmd/handlers` to differentiate between different error cases and
situations.

### pkg

The purpose of this folder is to copy-paste the commonly used packages that get used often and yet, have no dependency
to the project and can be plugged in or out based on various needs.

## Ent

### Generate Ent

```shell
go generate ./internal/ent
```

### Generate new schema file

```shell
go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema <Name>
```