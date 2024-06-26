# JWT authentication in Go

Учебный проект, направленный на практическое освоение разработки сервиса для регистрации и аутентификации пользователей. В рамках этого проекта используются технологии PostgreSQL для хранения данных, фреймворк Gin для создания веб-приложений на языке Go и инструменты для миграции базы данных. Все компоненты приложения запускаются в контейнерах Docker, обеспечивая удобную и переносимую среду разработки.

---

## Команды Docker

- **docker-build:** Сборка контейнеров Docker.
```bash
make docker-build
```
- **docker-up:** Запуск контейнеров Docker.
```bash
make docker-up
```
- **docker-down:** Остановка контейнеров Docker.
```bash
make docker-down
```

## Команды миграции базы данных

- **migrate-up:** Применение миграций базы данных вверх.
```bash
make migrate-up
```
- **migrate-down:** Откат миграций базы данных.
```bash
make migrate-down
```
- **migrate-create:** Создание новой миграции.
```bash
make migrate-create
```

Пожалуйста, убедитесь, что перед использованием команды `migrate-up` и `migrate-down` база данных PostgreSQL запущена и доступна по адресу `localhost:5432`, а также у вас установлен инструмент `migrate`.

---

## Маршруты API

### Получение токена доступа

**Метод:** POST  
**Путь:** `/api/token`  
**Описание:** Этот маршрут используется для получения токена доступа.

### Регистрация пользователя

**Метод:** POST  
**Путь:** `/api/user/register`  
**Описание:** Этот маршрут используется для регистрации нового пользователя.

### Проверка безопасности

**Метод:** GET  
**Путь:** `/api/secured/ping`  
**Описание:** Этот маршрут используется для проверки безопасности и доступа к защищенным ресурсам.