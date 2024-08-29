#!/bin/bash

# Добавление заметки
curl -X POST -H "Content-Type: application/json" -d '{"id":"1", "content":"This is a test note"}' http://localhost:8080/notes

# Получение списка заметок
curl -X GET http://localhost:8080/notes