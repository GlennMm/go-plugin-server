package engine

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"plugin"
	"time"

	"web_api_engine/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Engine struct {
	Router    *mux.Router
	ApiRouter *mux.Router
	DbEngine  *gorm.DB
	Ctx       context.Context
}

func NewEngine() *Engine {
	router := *mux.NewRouter()
	router.Use(middlewares.CorsMiddleware)
	api_router := router.PathPrefix("/api/").Subrouter()

	return &Engine{
		Router:    &router,
		ApiRouter: api_router,
		Ctx:       context.Background(),
	}
}

func (e *Engine) Run() {
	err := e.init()
	if err != nil {
		fmt.Println(err)

		panic("Failed to initialized the engine.")
	}
	err = e.loadPlugins()
	if err != nil {
		fmt.Println(err)

		panic("Failed to load all plugins.")
	}
	e.startServer()
}

func (e *Engine) init() error {
	// FIX: get DB name from .env file
	db_engine, err := gorm.Open(sqlite.Open("app_db.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		// return err
	}
	e.DbEngine = db_engine
	e.Ctx = context.WithValue(e.Ctx, "my_store", db_engine)

	// TODO: load middlewares
	fmt.Println("engine has been warmed up.")
	return nil
}

func (e *Engine) loadPlugins() error {
	e.Router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}
	files, err := filepath.Glob(filepath.Join(dir, "/modules/*.so"))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if len(files) == 0 {
		fmt.Println("No plugins found.")
		return nil
	}

	for _, file := range files {

		plug, err := plugin.Open(file)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Get the symbol named "Greeter" from the plugin.
		Load, err := plug.Lookup("Load")
		if err != nil {
			fmt.Println(err)
			return err
		}

		Load.(func(*mux.Router, *gorm.DB))(e.ApiRouter, e.DbEngine)

	}
	fmt.Printf("%d plugin(s) loaded.\n", len(files))
	return nil
}

func (e *Engine) startServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: "0.0.0.0:5000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      e.Router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("Starting the server at " + srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
