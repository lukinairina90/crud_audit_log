# Сервис логирования событий для CRUD APP
## crud_movies

### Сначала: 
#### ```go get google.golang.org/grpc```
#### ```go get google.golang.org/protobuf```

### Потом описываем audit.proto и генери

### генерим audit.pb.go:
#### ```protoc --go_out=pkg --go-grpc_out=pkg proto/audit.proto```


### Потом устанавливаем драйвер монги:
#### ```go get go.mongodb.org/mongo-driver/mongo```

### Для конфигов скачиваем библиотеку:
#### ```go get github.com/kelseyhightower/envconfig```
### и делаем импорт в файле config.go:
#### ```import "github.com/kelseyhightower/envconfig"```

### Запуск
Для запуска необходимо указать переменные окружения, например в файле .env

```
export DB_URI=mongodb://localhost:27017
export DB_USERNAME=admin
export DB_PASSWORD=g0langn1nja
export DB_DATABASE=audit
export SERVER_PORT=9000
```

Сборка и запуск
```
source .env && go build -o app cmd/main.go && ./app
```

Для mongo можно использовать Docker

```
docker run --rm -d --name audit-log-mongo \
> -e MONGO_INITDB_ROOT_USERNAME=admin \
> -e MONGO_INITDB_ROOT_PASSWORD=g0langn1nja \
> -p 27017:27017 mongo:latest
```





