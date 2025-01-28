# Используем официальный Go образ
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем исходный код
COPY . .

# Строим приложение
RUN go build -o main ./cmd/main.go

# Открываем порт для API
EXPOSE ${APP_PORT}

# Запускаем приложение
CMD ["./main"]
