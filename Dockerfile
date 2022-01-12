FROM golang:1.17-alpine

#where data is stored
WORKDIR /app

#recreate current working directory and cd to it
COPY . .

RUN go build -o main /app/cmd/main-server/main.go

EXPOSE 4000

CMD [ "/app/main" ]
