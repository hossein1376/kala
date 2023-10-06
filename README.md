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

(For older iterations and alternative structures, they can be found on other branches)

### In a look

```
.
├── cmd
│   └── api
│      ├── main.go
│      ├── run.go
│      └── ...
├── config
│   └──  state.go
├── internal
│   ├── data
│   │   ├── Models.go
│   │   ├── Mocks.go
│   │   └── ...
│   ├── ent
│   ├── handlers
│   │   ├── routes.go
│   │   ├── handlers.go
│   │   └── ...
│   ├── structure
│   └── transfer
│       ├── errors.go
│       └── response.go
├── pkg
│   ├── Json
│   ├── Logger
│   └── ...
├── go.mod
└── ...
```

### cmd/api

This module is tasked with retrieving the configurations, their validation, opening database connections and then
finally starting the server.  
It also handles the graceful shutdown as well.

The configurations, logger and an instance of `Model` will be passed down from here to the `handlers` module.

### config

The configuration and setting structures are defined here, as they'll be used inside application to control the API's
behaviour.

### internal/data

Database queries and transactions will take place here. Each table features a struct with db methods attached to it,
they all will be initialized with the `Models` struct by a database connection.  
With the use of interfaces, it's easy to write mocks for test cases.

### internal/handlers

`handlers` module feature the struct `handler` which all handlers are a receiver function to it. It has a single
exported receiver function named `Router` which will be called inside the `cmd/api/run.go` to instantiate the router.

Multiple structs such as `Json` and `Response` will be embedded inside the `handler` struct so they can be used directly
by the handlers.

### internal/structure

This module consists of structs to define the request, response and data structure; hence the name.  
It's equivalent of `dto` in some other architectures.

### internal/transfer

This module is tasked with the data transfer inside the application, as well as to the clients. It features a general
HTTP response function, besides many other helper functions each meant for a specific status code.

`transfer` also is home to handful of structs that all implement `error` interface. They are meant to be used as a mean
of conveying state between other layers and the handlers.

### pkg

The purpose of this folder is to copy-paste the commonly used packages that get used often and yet, have no dependency
to the project and can be plugged in or out based on various needs. Packages such as `Json` and `Logger`, among others,
reside here.

## Running

Install dependencies:

```shell
go mod tidy
```

Then, run the application:

```shell
go run ./cmd/api 
```

Pass the configurations as flag arguments, or provide them as environmental variables.

To run in development mode, add the following flag: `-env dev`

## Ent

### Generate Ent

```shell
go generate ./internal/ent
```

### Generate new schema file

```shell
go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema <Name>
```

## TODOs:

- [x] Simplify the application logic
- [x] ~~Improve error handling~~ Created new `transfer` module
- [x] Load configurations from a file as an option
- [x] Defining the log-level from the configurations
- [ ] Add more logs to the application
- [ ] Add a new layer between the `handlers` and `data` packages
- [ ] Integrate Websocket