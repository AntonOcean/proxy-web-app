### Как установить докер
```sudo apt install docker.io``` <br><br>
```sudo systemctl start docker``` <br><br>
```sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose```<br><br>
```sudo chmod +x /usr/local/bin/docker-compose```<br><br>

### Запуск на машине
```sudo docker-compose build web``` <br><br>
```sudo docker-compose up web``` <br><br>

### Сервисы
```127.0.0.1:8080 - api для сохранения и повтора запросов``` <br><br>
```127.0.0.1:9090 - proxy сервер``` <br><br>

### API
```127.0.0.1:8080/repeat/{id} - выполнить запрос с id``` <br><br>
```127.0.0.1:8080/all - посмотреть все запросы``` <br><br>