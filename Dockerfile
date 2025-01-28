# Используем официальный Go образ
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем необходимые зависимости для Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем исходный код
COPY . .

# Генерируем Swagger документацию
RUN swag init -g cmd/main.go

# Строим приложение
RUN go build -o main /app/cmd/main.go

# Открываем порт для API
EXPOSE ${APP_PORT}

# Запускаем приложение
CMD ["./main"]
