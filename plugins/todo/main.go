package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"todo/handlers"
	"todo/models"

	"github.com/gorilla/mux"
	"xorm.io/xorm"
)

type TodoPlugin struct {
	Name string
}

// addMiddleWare implements interfaces.IPlugin.
func (*TodoPlugin) addMiddleWare(r *mux.Router, db *xorm.Engine) error {
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "store", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	return nil
}

// migrateModels implements interfaces.IPlugin.
func (*TodoPlugin) migrateModels(db *xorm.Engine) error {
	if err := db.Sync(new(models.Todo)); err != nil {
		return err
	}

	return nil
}

// registerRoutes implements interfaces.IPlugin.
func (p *TodoPlugin) registerRoutes(r *mux.Router, db *xorm.Engine) error {
	router := r.PathPrefix(fmt.Sprintf("/" + p.Name)).Subrouter()

	router.HandleFunc("", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"message": "Todo Web Api Module."})
	})

	LoadRoutes(router)

	return nil
}

func LoadRoutes(r *mux.Router) {
	r.HandleFunc("/new", handlers.NewTodo).Methods("POST", "OPTIONS")
	r.HandleFunc("/{id:[0-9]+}", handlers.ReadOneTodo).Methods("GET", "OPTIONS")
	r.HandleFunc("/{id:[0-9]+}", handlers.DeleteTodo).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/all", handlers.ReadAllTodo).Methods("GET", "OPTIONS")
	r.HandleFunc("/{id:[0-9]+}", handlers.UpdateTodo).Methods("PUT", "PATCH", "OPTIONS")
}

func (p *TodoPlugin) RegisterPlugin(r *mux.Router, db *xorm.Engine) error {
	if err := p.migrateModels(db); err != nil {
		return err
	}
	if err := p.addMiddleWare(r, db); err != nil {
		return err
	}
	if err := p.registerRoutes(r, db); err != nil {
		return err
	}
	return nil
}

func CreateTodoPlugin(Name string) *TodoPlugin {
	return &TodoPlugin{
		Name,
	}
}

func Load(router *mux.Router, db *xorm.Engine) {
	plugin := CreateTodoPlugin("todo")
	plugin.RegisterPlugin(router, db)
}
