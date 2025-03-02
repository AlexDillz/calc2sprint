# syntax=docker/dockerfile:1

###############################
# Stage 1: Builder
###############################
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем бинарники для сервера и агента
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o agent ./cmd/agent

###############################
# Stage 2: Server Image
###############################
FROM alpine:latest AS server_image

WORKDIR /app
COPY --from=builder /app/server .

# Открываем порт, указанный в конфигурации (по умолчанию 8080)
EXPOSE 8080

# Запуск серверного приложения
CMD ["./server"]

###############################
# Stage 3: Agent Image
###############################
FROM alpine:latest AS agent_image

WORKDIR /app
COPY --from=builder /app/agent .

# Запуск агента
CMD ["./agent"]
