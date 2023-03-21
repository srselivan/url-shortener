### Запуск с postgres
* Если используется `docker`:
```dockerfile
CMD ["bin/app", "-d"]
```
Или с помощью переменной окружения `USE_POSTGRES` равной `true`
Замечание: `USE_POSTGRES` перекрывает флаг `-d`
* Иначе с помощью `-d` при запуске или `USE_POSTGRES=true`

При запуске не из контейнера изменить `PG_HOST` (или удалить, 
по умолчанию будет устанавливаться `localhost`) 
### Примеры использования (`bash`)
POST
```
curl --location 'http://localhost:8080/' \
--header 'Content-Type: text/plain' \
--data 'https://google.com'
```
Response body:
```
http://localhost:8080/h
```
GET
```
curl --location 'http://localhost:8080/h'
```
Response body:
```
https://google.com
```