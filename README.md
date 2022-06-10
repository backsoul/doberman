### START
```
 docker-compose up -d
```
### MIGRATE 
```
docker exec -it doberman sh
go run cmd/doberman/migrate/main.go
```
### TESTS
```
docker exec -it doberman sh
go test ./... -v
```

### RABBIT MQ
```
http://0.0.0.0:15672/
user: guess
password: guess
```
### CONSUMER 
```
go run cmd/doberman/queue/consumer.go 
```
### VIEW MESSAGES CONSUMER
```
docker logs -f doberman-consumer
```
### QUEUE
```
doberman.CreateArticle
```

### ENDPOINT 
```
http://localhost:3000/hello?message=hola
```
