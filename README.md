# Subscription Service

REST-сервис для агрегации данных об онлайн-подписках пользователей.

Тестовое задание Junior Golang Developer.

## Возможности

* CRUD операции над подписками
* Подсчет суммарной стоимости подписок за период
* Фильтрация по пользователю и названию сервиса
* PostgreSQL + миграции
* Конфигурация через `.env`
* Логирование
* Swagger документация
* Docker Compose запуск

---

## Стек технологий

* Go (Gin)
* PostgreSQL
* pgx
* logrus
* Swagger (swaggo)
* Docker / Docker Compose

---

## Структура проекта

```
subscription-service/
├── cmd/server          # точка входа
├── internal/
│   ├── config          # конфигурация
│   ├── db              # подключение к БД
│   ├── handler         # HTTP обработчики
│   ├── model           # модели
│   ├── repository      # доступ к БД
├── migrations          # SQL миграции
├── swagger             # swagger файлы
├── dockerfile
├── docker-compose.yml
└── .env
```

---

## Запуск проекта

### 1. Клонирование

```
git clone <repo_url>
cd subscription-service
```

### 2. Запуск через Docker

```
docker-compose up --build
```

Сервис будет доступен:

```
http://localhost:8080
```

---

## Swagger документация

```
http://localhost:8080/swagger/index.html
```

Через Swagger можно тестировать API.

---

## API

### Создание подписки

POST `/subscriptions`

Пример:

```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "2025-07-01"
}
```

---

### Получение подписок

GET `/subscriptions`

Фильтры:

* user_id
* service_name

---

### Обновление подписки

PUT `/subscriptions/{id}`

---

### Удаление подписки

DELETE `/subscriptions/{id}`

---

### Суммарная стоимость

GET `/subscriptions/summary`

Параметры:

* user_id
* service_name
* from (MM-YYYY)
* to (MM-YYYY)

Ответ:

```json
{
  "total": 400
}
```

---

## Конфигурация

Файл `.env`:

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=subscriptions
PORT=8080
```

---

## Миграции

SQL файлы находятся в папке `migrations/`.

Применяются автоматически при запуске Docker.

---

## Логирование

Используется logrus.

Логи выводятся в консоль контейнера.

---

## Автор

Test task implementation.
