# Тестовое задание для backend-разработчика

Необходимо реализовать API на Golang+Gin+MySQL или PostgreSQL на выбор для ToDo-доски из этого репозитория.

# Необходимые для реализации методы

`GET /todos` - Список всех записей. Ответ в виде массива объектов, структура объекта:

`{
	id: number;
    title: string;
    completed = false;
    dueDate = new Date();
}`

`GET /todos/:id` - Выбор отдельно взятой записи

`POST /todos` - Создание записи. Ответ от сервера вида

`{
	id: number;
    title: string;
    completed = false;
    dueDate = new Date();
}`

`PUT /todos/:id` - Обновление записи

`DELETE /todos/:id` - Удаление записи

# Предварительная подготовка

Для запуска проекта должен быть установлен Node.js

Устанавливаем Angular CLI

`npm install -g @angular/cli`

Переходим в папку с проектом и устанавливаем его

`npm install`

## Development server

Для запуска проекта в dev-режиме -  `ng serve` . Сам проект будет доступен по адресу `http://localhost:4200/`. 

## Build

Для сборки проекта выполнить `ng build` . Собранный дистрибутив будет находиться в папке `dist/` в корне проекта. Сборку для размещения на хостинге следует выполнять с ключом `--prod`


