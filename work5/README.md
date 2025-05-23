# CRUD REST API для ресурса `users`

## Первичная настройка
- `copy .env.sample .env` - инициализация файла настроек переменных окружения
- `docker-compose up -d` - запуск внешних сервисов
- `goose up` - инициализация базы данных, запуск миграций


## Пример использования
- `go run cmd/app/main.go` - запуск приложения
- `go test -tags imtegration ./...` - запуск интеграционных тестов (осуществлять после запуска основного приложения)
- `go test ./...` - запуск usecase тестов
- `go run cmd/http/main.go` - запуск http клиента (поддерживаемые команды: create, get, del, list, help)
- `go run cmd/grpc/main.go` - запуск grpc клиента (поддерживаемые команды: create, get, del, list, help)
