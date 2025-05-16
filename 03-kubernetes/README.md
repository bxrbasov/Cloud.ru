## Простое web-приложение в Kubernetes

Простое web-приложение, разворачиваемое в Kubernetes, состоит из двух компонентов:

+ Ingress Nginx для маршрутизации внешнего трафика;
+ Веб-сервер на основе Nginx.

Приложение доступно извне кластера через Ingress.

Настройка:

1. В первую очередь следует проверить minikube addons:
```
minikube addons list
``` 
2. В случае, если ingress отключен:
```
minikube addons enable ingress
```
3. Для подключения по хосту ```simple.web.app``` необходимо получить ```minikube ip```, затем перейти в file ```hosts``` и настроить локальный DNS (например, ```192.168.0.254 hello.local```):

    + Для Linux расположение: ```/etc/hosts```
    + Для Windows расположение: ```C:\Windows\System32\drivers\etc\hosts```

4. Для просмотра веб-страницы перейти на "simple.web.app" в браузере.