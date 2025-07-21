# 1. Используем официальный образ Golang
FROM golang:1.24.5

# 2. Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# 3. Копируем go.mod и go.sum и качаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# 4. Копируем весь проект
COPY . .

# 5. Компилируем приложение
RUN go build -o server ./cmd/server

# 6. Указываем команду, которая будет запущена при старте контейнера
CMD ["./server"]
