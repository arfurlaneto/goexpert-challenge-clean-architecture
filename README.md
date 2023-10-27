# goexpert-challenge-clean-architecture

## Running the application:

Start `MySQL` and `RabbitMQ`:
```bash
docker-compose up -d 
```

With the database up, run database migrations:
```bash
make run-migrations
```

Start the application:
```
make run
```

## Querying

### HTTP

You can use the sample requests from [create_order.http](./api/create_order.http):
```bash
curl -d '{"id":"1","price": 100.5,"tax":0.5}' -H "Content-Type: application/json" -X POST http://localhost:8000/order 

curl http://localhost:8000/order 
```

### gRPC

You can use evans (https://github.com/ktr0731/evans):

```bash
evans -r repl
package pb
service OrderService
call CreateOrder
call ListOrders
```

### GraphQL

You can use the playground running at http://localhost:8080/ with the following queries:

```graphql
mutation createOrder {
  createOrder(input: { id: "1", price: 10, tax: 0.5 }) {
    id
    price
    tax
    finalPrice
  }
}

query listOrders {
  orders {
    id
    price
    tax
    finalPrice
 	}
}
```

# Tools

https://github.com/google/wire  
https://grpc.io/docs/languages/go  
https://github.com/99designs/gqlgen  
