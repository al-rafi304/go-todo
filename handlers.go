package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func completeTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true

			jsonResponse(w, http.StatusOK, map[string]string{
				"message": fmt.Sprintf("Completed - %s", todo.Title),
			})
			return
		}
	}

	jsonResponse(w, http.StatusNotFound, map[string]string{
		"error": "Todo not found",
	})

}
