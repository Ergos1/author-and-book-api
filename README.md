# Домашнее задание 4

## Запуск
Запустить коммандер:
``make run``

## Пример:
Получить автора с айди 1:
``make run ARGS="get-author --author_id=1"``

Создать автора с айди 1:
``make run ARGS="create-author --author_id=3 --author_name=Yera"``


## Шаги
1) Внедрить в проект функционал консольных команд так, чтобы при добавлении новой команды понадобилось минимум изменений
2) Логика команд должна быть изолирована и в то же время работать в рамках одного шаблона
3) Добавить команду help, которая выводит информацию обо всех имеющихся консольных командах
4) Добавить команду spell с единственным аргументом - словом на английском языке
5) Команда spell принимает на вход слово, а по результатам работы выводит в консоль все буквы этого слова через пробел

## Дополнительно
💎 Реализовать команду - сильно упрощенный gofmt. На вход принимает *.txt файл, на выходе перед каждым абзацем вставляет таб и ставит точку в конце предложений.
# Домашнее задание 3

## Запуск:
Запустить приложение:

``make app_run``

Запутстиь миграцию:

``make migration-up``

.env пример:

POSTGRES_DB=test

POSTGRES_USER=test

POSTGRES_PASSWORD=test

POSTGRES_PORT=5433

POSTGRES_HOST=localhost

SERVER_ADDRESS=:9000

## Пример:
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

Обновить книгу:

``curl -X PUT -i http://localhost:9000/books/1 -d '{"id": 1, "name": "Yera", "rating": 1, "author_id": 1}'``

Удалить книгу:

``curl -X DELETE -i http://localhost:9000/books/1``

## Дано:

1) Написать CRUD операции для работы с бд
   должны быть реализованы методы
   - GetByID
   - Update
   - Create
   - Delete
2) Написать миграции для создания таблиц используя goose
3) Разработать HTTP сервер. Сервер должен уметь работать с бд и релизовывать CRUD операции.
    - Метод **GetByID**.
        - В Query-параметрах принимать идентификатор ?id= и возвращать данные из бд.
        - если передан идентификатор, который отсутствует в бд, возвращать HTTP-статус 404
        - если Query-параметрах отсутсвует ?id=, возвращать HTTP-статус 400
        - если произошла внутренняя ошибка сервера, возвращать HTTP-статус 500
        - если запрос успешно выполнен, возвращать HTTP-статус 200 и данные в теле ответа

    - Метод **Create**
        - В теле запроса принимать идентификатор и данные: и добавлять в бд
        - если такой идентификатор уже существует, возвращать HTTP-статус 409
        - если произошла внутренняя ошибка сервера, возвращать HTTP-статус 500
        - если запрос успешно выполнен, возвращать HTTP-статус 200

    - Метод **Delete**
        - В Query-параметрах принимать идентификатор ?id= и удалять из бд
        - если такого идентификатора не существует, возвращать HTTP-статус 404
        - если произошла внутренняя ошибка сервера, возвращать HTTP-статус 500
        - если запрос успешно выполнен, возвращать HTTP-статус 200

    - Метод **Update**
        - В теле запроса принимать идентификатор и данные и обновлять бд по ключу
        - если такого идентификатора не существует, возвращать HTTP-статус 404
        - если произошла внутренняя ошибка сервера, возвращать HTTP-статус 500
        - если запрос успешно выполнен, возвращать HTTP-статус 200

В случае ошибки в тело ответа пишем текст ошибки

4) Запустить сервер на порту 9000



💎 Внедрить доп сущность, которая будет связана с базовой таблицей. 
Минимум 1 доп поле кроме главного и внешнего ключей. Сущность должна отдаваться в ручках из пункта 2, будучи логически связанной с основной сущностью. Пример: главная сущность - пост, доп сущность-коммент. Если в Get отдается пост и все его комментарии это зачет. Если при удалении поста каскадно удаляются все комментарии тоже зачет. Покрывать все ручки не обязательно  

💎 В ридми приложить curl запросы, на каждую ручку. Запросы должны быть валидными и возвращать 200

Ограничения дз:
Нельзя использовать orm или sql билдеры
Для реализации http сервера можно использовать как net/http так и gin/fasthttp и прочее
Предметную область выбрать самостоятельно. Можно использовать пост/комментарий


