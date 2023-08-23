package interfaces

import "github.com/gorilla/mux"

type IPlugin interface {
	migrateModels() error
	addMiddleWare() error
	registerRoutes() error
	RegisterPlugin(r *mux.Router) error
}
