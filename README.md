# Тестовое задание для стажёра Backend
## Сервис динамического сегментирования пользователей
Микросервис для сегментирования пользователей на тестовые группы

Используемые технологии:

GIN (веб фреймворк)
PostgreSQL (в качестве хранилища данных)
Docker (для запуска сервиса)
Swagger (для документации API)
GORM (ORM библиотека)
# Usage
Запустить сервис можно с помощью команды `sh docker-compose up --build -d`

Документацию после завпуска сервиса можно посмотреть по адресу http://localhost:8080/swagger/index.html с портом 8080 по умолчанию

Для запуска тестов необходимо выполнить команду make test, для запуска тестов с покрытием make cover и make cover-html для получения отчёта в html формате

Для запуска линтера необходимо выполнить команду make linter-golangci