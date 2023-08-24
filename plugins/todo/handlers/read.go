package handlers

import (
	"fmt"
	"net/http"
	"web_api_engine/utils"

	"todo/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ReadOneTodo(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("store").(*gorm.DB)
	if !ok {
		fmt.Println("Db is null")
	}

	vars := mux.Vars(r)

	id := vars["id"]

	var todo models.Todo
	result := db.Where("id = ?", id).Find(&todo)
	if result.Error != nil {
		errs := []string{fmt.Sprintf("Todo with id %s was not found", id), result.Error.Error()}
		utils.Respond[interface{}](w, nil, errs, http.StatusNotFound)
		return
	}
	errs := []string{}
	utils.Respond[models.Todo](w, todo, errs, http.StatusOK)
}

func ReadAllTodo(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("store").(*gorm.DB)
	if !ok {
		fmt.Println("Db is null")
	}
	todos := []models.Todo{}
	result := db.Find(&todos)
	if result.Error != nil {
		fmt.Println(result.Error)
		errs := []string{result.Error.Error()}
		utils.Respond[interface{}](w, nil, errs, http.StatusNotFound)
		return
	}
	errs := []string{}
	utils.Respond[[]models.Todo](w, todos, errs, http.StatusOK)
}
