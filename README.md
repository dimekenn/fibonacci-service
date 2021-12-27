<h2>Сервис возвращающий массив чисел фибоначи</h2>

### Сервис имеет два api, GRPC и REST.

### Для работы сервиса необходим Docker, после в корневой папке сервиса пропишите:
```
docker-compose up --build fibonacci-service
```

#### REST api для вычисления будет доступен по адресу: 
```
http://localhost:8080/api/v1/fibonacci?x=0&y=7
```

#### для теста GRPC api нужно установить EVANS cli, прописать путь к прото файлам и порт 8081
```
evans /proto/fibonacci.proto -p 8081
```
#### и вызвать метод GetFibonacciSlice

```
call GetFibonacciSlice
```

#### где x - число откуда сервис должен начать для вичсления, y - последнее число вычисления.
