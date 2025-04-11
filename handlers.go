package main

import (
	"encoding/json"
	"net/http"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	todo.ID = nextId
	nextId++
	todos = append(todos, todo)

	jsonResponse(w, http.StatusOK, todo)
}
