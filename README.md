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
Запустить сервис можно с помощью команды ```docker-compose up --build -d```

Документация доступна по адресу http://localhost:8080/swagger/index.html
# Examples
Примеры запросов
### Создание сегмента 
Пример запроса:
```
curl --location --request POST 'http://localhost:8080/api/v1/segments' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"AVITO_TEST_VOICES"
}'
```
Пример ответа:
```json
{
    "name" : "AVITO_TEST_VOICES"
}
```
## Удаление сегмента
```
curl --location --request DELETE 'http://localhost:8080/api/v1/segments/AVITO_TEST_VOICES'
```
## Добавление и удаление сегментов пользователя 
Прмиер запроса:
```
curl --location --request POST 'http://localhost:8080/api/v1/users/1/segments' \
--header 'Content-Type: application/json' \
--data-raw '{
    "add":[{"name" : "AVITO_TEST_CUPONS", "delete_at" : "2024-10-02T15:00:00.05Z"}],
    "delete" : ["AVITO_TEST_VOICES"]
}'
```
Пример ответа:
```json
{   
    "id" : 1,
    "segments" : [{
        "name" : "AVITO_TEST_COUPONS",
        "delete_at" : "2024-10-02T15:00:00.05Z"
    }]
}
```
## Создание сегмента с автоматическим зполнением 
### Автоматически добавляет процент пользователей в сегмент

Пример запроса:
```
curl --location --request POST 'http://localhost:8080/api/v1/segments' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "AVITO_FREE_DELIVERY",
    "per_cent" : 33
}'
```
Пример ответа:
```json
{
    "name" : "AVITO_FREE_DELIVERY",
    "per_cent" : 33
}
```
## Получение отчета по пользователю за определенный период
Пример запроса:
```
curl --location --request GET 'http://localhost:8080/api/v1/users/1/segments/history?from=2007-10&to=2023-05'
```
Пример ответа:
```json
{
    "url" : "/files/reports/550e8400-e29b-41d4-a716-446655440000.csv"
}
```
