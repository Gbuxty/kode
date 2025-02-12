# Используем официальный образ Go
FROM golang:1.21.6-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Устанавливаем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o main .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]