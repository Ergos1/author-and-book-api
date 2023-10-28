# Homework 6

## Заметка



## Запуск

Запустить приложение:

``make run``

Запустить все зависимые сервисы(бд, кафка):

``make up-deps``

Завершить все зависимые сервисы(бд, кафка):

``make down-deps``

## Пример запросов:

Создать автора:

``curl -X POST -i http://localhost:9000/authors -d '{"id": 1, "name": "Yera"}'``

Получить автора:

``curl -X GET -i http://localhost:9000/authors/1``

Обновить автора:

``curl -X PUT -i http://localhost:9000/authors/1 -d '{"name": "Yera3"}'``

Создать кингу:

``curl -X POST -i http://localhost:9000/books -d '{"id": 1, "name": "Yera", "rating": 1, "author_id": 1}'``

Получить книгу:

``curl -X GET -i http://localhost:9000/books/1``

## Задание

Требования:
1) При CRUD операциях из домашнего задания 3 недели записываем события по методам в кафку
2) Событие должно хранить время возникновения, тип и сырой запрос
3) По ходу взаимодействия с приложением в консоль необходимо выводить данные по событиям из кафки 
4) Ивенты **нельзя** сразу выводить в консоль, минуя кафку
5) Решение должно быть покрыто unit тестами

## Дополнительно
💎 Написать интеграционные тесты на взаимодействия с кафкой

