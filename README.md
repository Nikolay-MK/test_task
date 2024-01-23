# Selltech

Этот проект представляет собой пример приложения на Go с использованием Docker.

## Сборка и запуск с Docker

```bash
docker-compose up
```

## ЗАПРОСЫ
Запрос для обновление таблицы
```
http://127.0.0.1:8080/update 
```

Запрос для получение состояния
```
http://127.0.0.1:8080/state
```

Получение списка всех возможных имён человека из локальной базы данных с указанием
основного uid в виде JSON
```
http://127.0.0.1:8080/get_names?name={SOME_VALUE}&type={strong|weak}
```