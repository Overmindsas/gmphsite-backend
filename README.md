# Приложение для работы с api (получение, запись в бд, вывод содержимого базы данных)

* Объявляется интерфейс DataService в пакете services/user.go
    * имплиментируется в services/user.impliment.go для работы с бд
    * имплиментируется в handlers/handler.go для работы со стороннем api
* Есть 2 маршрута (/getdataresp, /getalldata)
    * /getdataresp - делает запрос к api, получает json и записывает в бд (GET)
    * /getalldata - получает все данные из бд (GET)
* Используется база данных postgres, запущенная в контейнере


## Примеры запросов:

* 127.0.0.1:8080/getdararesp
* 127.0.0.1:8080/getalldata


## Команды для запуска

* go build cmd/main.go 
* sudo docker-compose up
* cd cmd -> ./main
