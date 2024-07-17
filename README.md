# Time Tracker API

## Содержание
- [Установка](#установка)
- [Настройка](#настройка)
- [Запуск](#запуск)
- [API документация](#api-документация)


## Установка
Клонируйте репозиторий:
git clone https://github.com/typaml/Effective_intern

Установите зависимости:
go mod tidy

## Настройка
Создайте файл .env в корневом каталоге проекта и добавьте в него необходимые переменные окружения:
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=timetracker
PORT=8080

## Запуск
Запустите сервер:
go run main.go

## API документация

После запуска сервера вы можете получить доступ к документации API по адресу:
http://localhost:8080/swagger/index.html
