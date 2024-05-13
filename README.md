<h1 align="center">
   <a href="#"> Clean Architecture with Golang </a>
</h1>

<h3 align="center">
   This Go program allows users to Create new service Orders by sendind Price and Tax, calculating the FinalPrice and List All created service Orders. This project was created following concepts of Clean Architecture (By Robert Cecil Martin).
</h3>

<h4 align="center"> 
	 Status: Finished
</h4>

# About
In this project we use REST, gRPC, GraphQL and RabbitMQ.

REST is a common architectural style for developing networked APIs. 
It utilizes HTTP methods (GET, POST, PUT, DELETE, etc.) to operate on resources, usually represented in JSON or XML format.

gRPC is a high-performance communication framework developed by Google. 
It enables communication between distributed services using Protocol Buffers (protobuf) to define the service interface and HTTP/2 for efficient binary communication

RabbitMQ is an open-source messaging system that implements the Advanced Message Queuing Protocol (AMQP). 
It enables asynchronous communication between distributed applications. 

GraphQL is a query language for APIs and a runtime for executing those queries with user-defined types. 
It provides clients with the ability to request exactly the data they need, nothing more and nothing less.

## Usage
1. Clone the repository.

## Dependencies
This program uses all dependencies in go.mod file
1. Run "go mod tidy" to download all dependencies.

#### Running Project

```bash

# Clone this repository
$ git clone https://github.com/renatafborges/go-clean-architecture.git

# Access folder contaning docker-compose.yaml and let's up all containers:
$ docker-compose up -d

# to up migration:
$ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

# in another terminal access cmd/ordersystem:
$ go run main.go wire_gen.go

# The server will start: REST, GRPC, GRAPHQL:
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080

# REST:
# in api directory you can create new orders or list all orders created
# just click on Send Request button in create_order.http file or list_order.http file.

# GRPC:
# In another terminal:
$ evans -r repl

# Select package
$ package pb

# Select service
$ service OrderService

# To Create a new Order
$ call CreateOrder

id (TYPE_STRING) => 500
price (TYPE_FLOAT) => 40
tax (TYPE_FLOAT) => 5

{
  "finalPrice": 45,
  "id": "300",
  "price": 40,
  "tax": 5
}

# Select service
$ service ListOrderService

# To List Categories
$ call ListOrder

# GRAPHQL:
# Open http://localhost:8080/

# To create a new Order:

mutation createOrder {
  createOrder(input:{id: "ddddd", Price: 12.2, Tax:2.0}) {
    id
    Price
    Tax
    FinalPrice
  }
}

# To list all Orders:

query queryOrders {
  orders {
    id
    Price
   	Tax
    FinalPrice
  }
}

# RABBITMQ:
http://localhost:15672/#/queues

## Notes
These are some of the key technologies you can use when developing Go applications to build RESTful APIs,
gRPC services, messaging systems with RabbitMQ and flexible APIs with GraphQL. 
Each has its own advantages and specific use cases, so choose the one that best fits your project's needs;

## Author
Made with love by Renata Borges 👋🏽 [Get in Touch!](Https://www.linkedin.com/in/renataborgestech)
