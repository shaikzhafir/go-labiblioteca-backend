trying out writing go backend with postgres db

using rancher desktop for local development

steps:

1. install rancher desktop
2. create a setenv.bat (Windows) file in the root folder of your project. fill with the env keys
3. Run setenv.bat in cmd to set ur env variables in environment
4. cd over to postgresql folder and run helm install postgres .
5. portforward the postgres service to connect from local go service
   1. `kubectl port-forward svc/postgres-postgresql 5433:5433`
6. Run 
   1. `go run cmd/main-server/main.go` 
7. Test CRUD at localhost:4000/books

## Structure of code  (starting from base to the router)

### Domain

- This is where the building blocks are defined. For this case, the base Book struct is defined here.

### Repository

- The layer that connects the application to the database layer.
- SQL queries are written here, keep to the simple CRUD for now

### Service

- This is where business logic for the application is defined.
- When there needs to be additional data modeling or massaging, its done here.
- Will be using methods defined in repository to get the data

### Handlers

- Where the routes are defined. Routes will be calling the different services depending on the use case
- Will be returning reponses to the user here and handling the http status codes here too

### How it all comes together

- In main.go, a handler is instantiated
- handler will need an instantiation of a service
- service will need an instantiation of a repository
- repository will need an instantiaton of a database connection, which is also done in main
- together they form a functional API

TODO

- tests
- logger
- better error handling
- authentication
- todo is listed in order of importance
- or not?
- find out next time on the price is right


## Protobuf stuff
1. Remove all pre existing generated protobuf go files
   
    `rm -rf pb/*.go`
2. Run this command to generate the go protobuf files based on the 
protobuf definitions 

    `protoc --proto_path=domain --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
 --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
 proto/*.proto domain/*.proto`