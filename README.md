# grpc-go

## Setup
```
grpc-go % docker-compose up -d
```

## Launch gRPC server
```
grpc-go % docker-compose exec app bash

root@1f3231f44621:/server# go run cmd/api/main.go
2022/04/03 14:44:06 Server is running!
```

## Call API
gRPC server remains up and running

```
ryo.naruse@dev grpc-go % docker-compose exec app bash
root@1f3231f44621:/server# go run client/main.go 
User is Tom 
```