trying out writing go backend with postgres db

using docker to setup db

steps:

1. install docker desktop and setup the shiz
2. setup the .env files as per your preferences, see sample (im just using cmd and running setenv.bat)
3. go to main directory and run docker-compose up -d to run postgres db
4. run 'go run main.go'
5. test at localhost:4000/books
