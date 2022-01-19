trying out writing go backend with postgres db

using docker to setup db

steps:

1. install docker desktop
2. create a setenv.bat (Windows) file in the root folder of your project. fill with the env keys
3. Run setenv.bat in cmd to set ur env variables in environment
4. Run docker build . -t my-images/biblioteca from same directory of Dockerfile. 
   This is to build the image to be used in docker-compose
4. Run docker-compose up -d 
5. Test at localhost:4000/books


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
