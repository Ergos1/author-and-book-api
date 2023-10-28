# Homework 6

## Заметка



## Запуск

Запустить приложение:

``make run``

Запустить все зависимые сервисы(бд, кафка):

``make up-deps``

Завершить все зависимые сервисы(бд, кафка):

``make down-deps``

Запустить миграции:

``make migration-up``

Откатить миграции:

``make migration-down``

Статус миграции:

``make migration-status``

Запуск юнит тестов:

``make test-unit``

Запуск интеграционных тестов:

``make test-integration``

## Пример запросов:

Создать автора:

``curl -X POST -i http://localhost:9000/authors -d '{"id": 1, "name": "Yera"}'``

Получить автора:

``curl -X GET -i http://localhost:9000/authors/1``

Обновить автора:

``curl -X PUT -i http://localhost:9000/authors/1 -d '{"name": "Yera3"}'``

Удалить автора:

``curl -X DELETE -i http://localhost:9000/authors/1``

Создать кингу:

``curl -X POST -i http://localhost:9000/books -d '{"id": 1, "name": "Yera", "rating": 1, "author_id": 1}'``

Получить книгу:

``curl -X GET -i http://localhost:9000/books/1``

## Пример .env:

POSTGRES_DB=test
POSTGRES_USER=test
POSTGRES_PASSWORD=test
POSTGRES_PORT=5433
POSTGRES_HOST=localhost
SERVER_ADDRESS=:9000

## Задание

Требования:
1) При CRUD операциях из домашнего задания 3 недели записываем события по методам в кафку
2) Событие должно хранить время возникновения, тип и сырой запрос
3) По ходу взаимодействия с приложением в консоль необходимо выводить данные по событиям из кафки 
4) Ивенты **нельзя** сразу выводить в консоль, минуя кафку
5) Решение должно быть покрыто unit тестами

## Дополнительно
💎 Написать интеграционные тесты на взаимодействия с кафкой

