## Команды, которые использовались для пуша образа в регистр

Для этого необходимо провести следующие действия:

1. Собрать Docker-образ из Dockerfile: 
```
docker build -t web_app .
``` 
где "-t web_app" - задаёт имя образа web_app.

2. Создать Docker-образ, который будет пушиться в приватный регистр докер хаб:
```
docker tag f9900df759cb bxrbasov/web_app:1.0
```
3. Залогиниться в докер хаб:
```
docker login
```
4. Запушить имеющийся Docker-образ в докер хаб:
```
docker push bxrbasov/web_app:1.0
```

В результате получим данный образ в Docker hub, остаётся только сделать его приватным.