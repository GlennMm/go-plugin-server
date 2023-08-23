package handlers

import (
	"fmt"
	"net/http"
	"web_api_engine/utils"

	"todo/models"

	"github.com/gorilla/mux"
	"xorm.io/xorm"
)

func ReadOneTodo(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("store").(*xorm.Engine)
	if !ok {
		fmt.Println("Db is null")
	}

	vars := mux.Vars(r)

	id := vars["id"]

	var todo models.Todo
	err := db.Where("id = ?", id).Find(todo)
	if err != nil {
		errs := []string{fmt.Sprintf("Todo with id %s was not found", id), err.Error()}
		utils.Respond[interface{}](w, nil, errs, http.StatusNotFound)
		return
	}
	errs := []string{}
	utils.Respond[models.Todo](w, todo, errs, http.StatusOK)
}

func ReadAllTodo(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("store").(*xorm.Engine)
	if !ok {
		fmt.Println("Db is null")
	}
	todos := []models.Todo{}
	err := db.Find(todos)
	if err != nil {
		fmt.Println(err)
		errs := []string{err.Error()}
		utils.Respond[interface{}](w, nil, errs, http.StatusNotFound)
		return

	}
	errs := []string{}
	utils.Respond[[]models.Todo](w, todos, errs, http.StatusOK)
}
