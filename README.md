# Golang API

## Rest API with Golang and MongoDB

At the moment need to have installed `Go` and `MongoDB` to use this code

See the section `Running the app with docker-compose`

This `API` try to get the best price (purchase and sale) of currencies from multiple `Operators`

### To start:

1. Clone this repo using `git clone https://github.com/wchopite/golang-mongo-api.git`
2. Rename the file `.env.example` to `.env` and setup the environment vars
6. Run the application with `go run main.go`
7. Access `http://localhost:1337/api/currency`

Each request to the API must include the following headers:

```
Content-Type: "application/json"
```

## Things to do after configuration and running

1. Create a `currency price` using the `POST /api/currency` endpoint and the next payload (for example):

```json
{
    "coin" : "USD",
    "purchase_value" : 60.58,
    "sale_value" : 58.59,
    "operator" : "Santander"
}
```

2. Then, you can use the other endpoint `GET /api/currency`. This endpoint return (for example):

```json
[
    {
        "Coin": "EUR",
        "Purchase_value": 65,
        "Sale_value": 69,
        "Operator": "Banco Nacion"
    },
    {
        "Coin": "USD",
        "Purchase_value": 61.5,
        "Sale_value": 57.5,
        "Operator": "Banco Nacion"
    }
]
```

You can get the best value only for one `currency` using in the `query params` the var `key` (this var represent the currency). For example:

`GET /api/currency?key=USD`

The return value:

```json
[
    {
        "Coin": "USD",
        "Purchase_value": 61.5,
        "Sale_value": 57.5,
        "Operator": "Banco Nacion"
    }
]
```
## Running the app with docker-compose

`Working in Progress`

1. `Install Docker`: https://docs.docker.com/compose/install/
2. `Install Docker Compose`: https://docs.docker.com/compose/install/

Once you have installed them, in the root folder you need to run `docker-compose up`. This command generate the containers with the `go app and mongo database`. Then you can test it using a Restful client, like Postman

## Tech Stack

1. Golang: https://golang.org/
2. Gorilla Mux: https://github.com/gorilla/mux
3. MongoDb: https://docs.mongodb.com/manual/
4. MongoDb Driver for go: https://docs.mongodb.com/ecosystem/drivers/go/
5. Godotenv: https://github.com/joho/godotenv

## To Do

1. Add `unit and integration test`
2. Add input validation on endpoints
3. Add `coverage`
4. Add a `cache layer` with `redis` for example
5. A better code organization. Divide in a better way the multiple layers
6. Add files configuration to `CI/CD` (codeship, google cloud build, etc)
7. Keep learning about `Go`, best practices, standards
8. Patterns, patterns, patterns...
