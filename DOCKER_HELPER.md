# DOCKER cmds

### Собрать контейнер
```shell
docker build -t my-golang-app .
```

### Запустить контейнер
```shell
docker run -p 8016:8080 --rm my-golang-app

docker compose up -d
```

### Остановить контейнер
```shell
docker stop (ID)
```

### Просмотр списка контейнеров
```shell
docker ps
```

### Установка docker на ubuntu
```shell

sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

```