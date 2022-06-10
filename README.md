### MIGRATE 
```
go run cmd/doberman/migrate/main.go
```
### TESTS
```
go test ./... -v
```

### RABBIT MQ
```
http://0.0.0.0:15672/
user: guess
password: guess
```
#### consumer 
```
go run cmd/doberman/queue/consumer.go 
```
### QUEUE
```
doberman.CreateArticle
```

### ENDPOINT 
```
http://localhost:3000/hello?message=hola
```
