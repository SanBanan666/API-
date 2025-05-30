# People API Service

Сервис для работы с данными о людях, включающий обогащение информации о возрасте, поле и национальности через внешние API.

## Описание проекта

Сервис предоставляет REST API для работы с данными о людях. При создании или обновлении записи о человеке, сервис автоматически обогащает данные дополнительной информацией:
- Возраст (через API agify.io)
- Пол (через API genderize.io)
- Национальность (через API nationalize.io)

## Технологии

- Go 1.21+
- PostgreSQL
- Gin Web Framework
- GORM (ORM для работы с базой данных)
- Swagger для документации API
- Docker (опционально)

## Требования

- Go 1.21 или выше
- PostgreSQL
- Docker (опционально)

## Структура проекта

```
.
├── main.go              # Точка входа приложения
├── models/             # Модели данных
│   └── person.go       # Модель Person
├── handlers/           # Обработчики HTTP запросов
│   └── person.go       # Обработчики для работы с людьми
├── services/           # Бизнес-логика
│   └── enrichment.go   # Сервис обогащения данных
├── database/           # Работа с базой данных
│   └── db.go          # Инициализация и конфигурация БД
├── migrations/         # Миграции базы данных
│   └── 001_init.sql   # Начальная миграция
└── .env               # Конфигурация окружения
```

## Установка

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd t3_juniorGo
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл .env на основе .env.example:
```bash
cp .env.example .env
```

4. Настройте переменные окружения в файле .env:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=people_db
SERVER_PORT=8080
```

## Запуск

1. Запустите PostgreSQL:
```bash
docker run --name people_db -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=people_db -p 5432:5432 -d postgres
```

2. Запустите приложение:
```bash
go run main.go
```

## API Endpoints

### Создание нового человека
- **URL**: `POST /api/people`
- **Тело запроса**:
```json
{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich" // опционально
}
```
- **Ответ**: Созданный человек с обогащенными данными

### Получение списка людей
- **URL**: `GET /api/people`
- **Параметры запроса**:
  - `page` (int, опционально) - номер страницы
  - `limit` (int, опционально) - количество записей на странице
  - `name` (string, опционально) - фильтр по имени
  - `surname` (string, опционально) - фильтр по фамилии
- **Ответ**: Список людей с пагинацией

### Обновление данных человека
- **URL**: `PUT /api/people/:id`
- **Тело запроса**: Аналогично созданию
- **Ответ**: Обновленные данные человека

### Удаление человека
- **URL**: `DELETE /api/people/:id`
- **Ответ**: 204 No Content

## Swagger документация

Доступна по адресу: http://localhost:8080/swagger/index.html

## Примеры запросов

### Создание нового человека
```bash
curl -X POST http://localhost:8080/api/people \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich"
  }'
```

### Получение списка людей
```bash
curl http://localhost:8080/api/people?page=1&limit=10
```

### Обновление данных человека
```bash
curl -X PUT http://localhost:8080/api/people/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich"
  }'
```

### Удаление человека
```bash
curl -X DELETE http://localhost:8080/api/people/1
```

## Разработка

### Добавление новых функций
1. Создайте новую ветку для разработки
2. Внесите изменения
3. Напишите тесты
4. Создайте pull request

### Тестирование
```bash
go test ./...
```

## Лицензия

MIT

## Автор

[Бабамуродов Мухаммаджон]
