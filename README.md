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


### GRPC SERVER

Запустить grpc server:

``make run-grpc-server``


### DATABASE

Поднять базу данных:

``make up-db``

Завершить базу данных:

``make down-db``

Запустить миграции:

``make migration-up``

Откатить миграции:

``make migration-down``

Статус миграции:

``make migration-status``


## Задание

Требования:
1) Переписать REST сервисы из домашнего задания 6 на gRPC 
2) Переписать логирование на структурное (с исользованием go.uber.org/zap)
3) Добавить трейсы

## Дополнительно
💎 Подключить gRPC-Gateway для обратной совместимости с REST
