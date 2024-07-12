# nerversitup-test

## code latest structure

.
└── Project/
├── cmd/  
│ ├── cmds/  
│ │ ├── rest.go  
│ │ └── root.go  
│ └── main.go
├── config/  
│ └── config.go
├── infrastructure/  
│ ├── mongodb.go
│ ├── redis.go
│ └── ...
├── internal/
│ ├── handlers/
│ │ ├── profilehandler.go
│ │ └── ...
│ ├── adapters/
│ │ ├── s3adapter.go
│ │ └── ...
│ ├── services/
│ │ ├── profile/
│ │ │ ├── profile.go
│ │ │ ├── service.g0
│ │ │ └── service_test.go
│ │ └── ...
│ ├── repositories/
│ │ ├── mongorepo/
│ │ │ ├── repositories.go
│ │ │ ├── profile.go
│ │ │ └── ...
│ │ └── ...
│ ├── ports/
│ │ ├── adapter.go
│ │ ├── service.go
│ │ ├── repository.go
│ │ └── ...
│ └── domain/
│ │ ├── adapter/
│ │ │ ├── s3.go
│ │ │ └── ...
│ │ ├── constant/
│ │ │ └── ...
│ │ ├── dto/
│ │ │ ├── profile.go
│ │ │ └── ...
│ │ └── model/
│ │ ├── profile.go
│ │ └── ...
├── middleware/
│ ├── error.go
│ └── ...
├── pkgs/
│ ├── utils.go
│ └── ...
├── protocol/
│ ├── init.go
│ ├── http.go
│ └── ...
└── .env

## code latest struct convention

- cmd : commands to run this app
- config : configuration mapping .env for use in this app
- infrastructure : code to new instant for use in service or anoter ...
- internal : core of code for this app
  - handlers : handler/controller
  - adapters : code for a proxy to port to external part
  - services : code for main business logic
  - repositories : code to query/update/delete data. no business logic here
  - ports: interface of all code
  - domain : all struct data
    - adapter : struct us in adapter
    - constant : a contstant
    - dto : data object transfer in/out
    - model : struct data in repository
- middleware : code to us as a middleware of app
- pkgs : internal package
- protocol : a protocol to use for app
- .env : environment
