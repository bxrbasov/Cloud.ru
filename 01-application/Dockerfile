# Используем официальный образ Go в качестве базового - https://hub.docker.com/_/golang
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходный код приложения
COPY . .

# Компилируем приложение
RUN go build -o main main.go

# Создаем минимальный образ для запуска приложения
FROM alpine:latest AS webapp

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем скомпилированный бинарный код из предыдущего этапа сборки
COPY --from=builder /app/main .

# Копируем шаблон index.html
COPY --from=builder /app/index.html .

# Устанавливаем права на исполнение
RUN chmod +x main

# Указываем порт, который прослушивает приложение
EXPOSE 8000

# Указываем команду, которую контейнер должен выполнить при запуске
ENTRYPOINT  ["./main"]