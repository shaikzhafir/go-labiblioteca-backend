trying out writing go backend with postgres db

using docker to setup db

steps:

1. install docker desktop
2. create a .env file in the root folder of your project. use the keys from sample and set your own values
3. while in root directory, run docker-compose up -d to run postgres db in docker
4. cd over to cmd/main-server and run 'go run main.go'
5. test at localhost:4000/books


## Structure of code  (starting from base to the router)
### Domain 
- This is where the building blocks are defined. For this case, the base 
Book struct is defined here.

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