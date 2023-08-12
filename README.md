# KALA

Kala (meaning goods) is an online shop backend written in Go.  
It aims to be fast, simple and extensible. As well as a showcase of my personal project structure.

## Project Structure

The general structure follows the guidelines I picked up
from [Let's Go Further](https://lets-go-further.alexedwards.net/),
but over time I found some shortcoming in that approach, which led me to explore for my desired structure.  
The following is the result of my experience of developing many Go projects, what I find more logical and extensible
approach,
compile-time checks (I'm looking at you, import cycle), my constant experiment with different designs and on top of all,
my personal taste.

While I find this overall structure to make sense the most, feel free to adapt it partially or completely for your own
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
|   ├── response
|
├── pkg
|   ├── json
|   ├── ...
|
├── main.go
├── go.mod
```

### cmd/api

Module `api` is tasked with retrieving the configurations, their validation, opening database connections and then finally
starting the server. It also handles the graceful shutdown as well.

The configurations, logger and an instance of `Model` will be passed down to `handlers`.

### cmd/handlers

`handlers` module feature the struct `handler` which all handlers are a receiver function to it. It has a single
exported receiver function named `Router` which will be called inside the `cmd/api/run.go` to instantiate the router.

Multiple structs such as `Json` and `Response` will be embedded inside the `handler` struct so they can be used directly by the handlers.

### cmd/state.go

This file consist of multiple struct to define the configurations and the "Application" data that will be passed down
to `handlers` from `cmd`.

The main reason to have this file here, instead of somewhere like `internal/structure`, is to avoid import cycle.

### internal/data

Database queries and transactions will take place here. Each table features a struct with db methods attached to it,
they all will be initialized with the `Models` struct by a database connection.  
With the use of interfaces, it's easy to write mocks for test cases.

### internal/structure

This module consists of structs to define the request, response and data structure, hence the name.

### internal/response

General response function (named `Respond`), as well as helper functions for specefic HTTP status codes reside here.  

Also custom defined errors that will be used to differentiate between different error cases and
situations are placed here; each has methods on them with varying degrees of information returned.

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
