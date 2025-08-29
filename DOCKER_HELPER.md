# DOCKER cmds

### Собрать контейнер
```shell
docker build -t my-golang-app .
```

### Запустить контейнер
```shell
docker run -p 8016:8080 --rm my-golang-app
```

### Остановить контейнер
```shell
docker stop (ID)
```

### Просмотр списка контейнеров
```shell
docker ps
```