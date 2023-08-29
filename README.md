# readbrains-project

This project consists of 2 separate golang services:
 - api-service-redbrains
 - crypto-service-redbrains



- crypto-service-redbrains is an rpc server that has only one method, [HashMovie](https://github.com/schiduluca/redbrains-project/blob/master/crypto-service-redbrains/proto/movie.proto#L7) 
- api-service-redbrains is an http server that has crypto-service-redbrains as dependency

### Steps to deploy the service
At the root of the project there's a `docker-compose.yaml` file which can be used to start the project locally.
To start the services just run the `docker-compose up` command.

The api-service-redbrains server will start on port `8082` but can be configured from docker-compose.yaml file via `PORT` env variable.

### Testing the services

To make an http request to api-service-redbrains service you can use curl

```
curl --location 'localhost:8082/api/v1/hash-movie-name' \
--header 'Content-Type: application/json' \
--data '{
"name": "test movie"
}'
```

Response should come in form of 
``{"Name":"ffc7bd296597df95f648ca8940e5ff5c097c2cb1e11b22d0cfd579dbdc6c6c59"}``
