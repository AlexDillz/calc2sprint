.PHONY: all build-server build-agent run-server run-agent run-all test clean

# По умолчанию собираем оба образа
all: build-server build-agent

# Сборка серверного образа
build-server:
	docker build --target server_image -t calc_server:server .

# Сборка образа агента
build-agent:
	docker build --target agent_image -t calc_server:agent .

# Запуск серверного контейнера
run-server:
	docker run --rm -p 8080:8080 calc_server:server

# Запуск контейнера агента
run-agent:
	docker run --rm calc_server:agent

# Запуск всей системы через docker-compose
run-all:
	docker-compose up --build

# Запуск всех тестов проекта
test:
	go test ./...

# Очистка артефактов сборки (если есть бинарные файлы)
clean:
	rm -f server agent