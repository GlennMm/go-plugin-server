package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web_api_engine/utils"

	"todo/models"

	"gorm.io/gorm"
)

func NewTodo(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("store").(*gorm.DB)
	if !ok {
		fmt.Println("Db is null")
		errs := []string{"No database instace"}
		utils.Respond[interface{}](w, nil, errs, http.StatusInternalServerError)
		return
	}
	todo := models.Todo{}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
		errs := []string{err.Error()}
		utils.Respond[interface{}](w, nil, errs, http.StatusNotFound)
		return

	}

	err = utils.DbInsert[models.Todo](db, &todo)
	if err != nil {
		fmt.Println(err)
		errs := []string{err.Error()}
		utils.Respond[interface{}](w, nil, errs, http.StatusNotFound)
		return
	}
	errs := []string{}
	utils.Respond[models.Todo](w, todo, errs, http.StatusOK)
}
