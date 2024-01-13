# Домашнее задание

## Запуск

### GRPC CLIENT:

Запустить grpc client:

``make run-grpc-client``

Посмотреть флаги для запуска:

``ARGS="--help" make run-grpc-client``

Создать автора:

``ARGS="-cmd=1 -id=1 -name=Example" make run-grpc-client``

Получить автора:

``ARGS="-cmd=0 -id=1" make run-grpc-client``

Обновить автора:

``ARGS="-cmd=2 -id=1 -name=Elpmaxe" make run-grpc-client``

Удалить автора:

``ARGS="-cmd=3 -id=1" make run-grpc-client``

### :gem: GRPC_GATEWAY

``make run-grpc-gateway``

### GRPC SERVER

Запустить grpc server:

``make run-grpc-server``


### DEPS

Поднять базу данных:

``make up-deps``

Завершить базу данных:

``make down-deps``

### DATABASE

Зайти в бд:

``docker exec -it hw3-test-db psql -U test``

Запустить миграции:

``make migration-up``

Откатить миграции:

``make migration-down``

Статус миграции:

``make migration-status``

## Jager

http://localhost:16686

## Шаги default

1. make up-deps
2. make run-grpc-server
3. Возможности "make run-grpc-client"

## Шаги :gem:

1. make up-deps
2. make run-grpc-gateway

Создать автора:

``curl -X POST -i http://localhost:9001/authors -d '{"id": 1, "name": "Yera"}'``

Получить автора:

``curl -X GET -i http://localhost:9001/authors/1``

Обновить автора:

``curl -X PUT -i http://localhost:9001/authors/1 -d '{"name": "Yera3"}'``

Удалить автора:

``curl -X DELETE -i http://localhost:9001/authors/1``
