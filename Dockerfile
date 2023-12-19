# Используем официальный образ Golang в качестве базового образа
FROM golang:1.21.0

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файл go.mod и go.sum и выполняем go mod download
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы из текущего каталога в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o main ./cmd

# Экспортируем порт, на котором работает приложение
EXPOSE 8080

# Устанавливаем переменную окружения config_path
ENV config_path /app/config/local.yaml

# Запускаем приложение
CMD ["./main"]