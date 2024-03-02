package gophermart

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/psfpro/gophermart/internal/gophermart/infrastructure/api/http/handler"
)

type Container struct {
	app *App
}

func (c Container) App() *App {
	return c.app
}

func NewContainer() *Container {
	config := NewConfig()

	// DB connection
	db, err := sql.Open("pgx", config.dsn)
	if err != nil {
		log.Printf("db open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("db connection error: %v", err)
	} else {
		log.Printf("db connection ok")
	}

	// HTTP handlers
	pingRequestHandler := handler.NewPingRequestHandler(db)
	notFoundHandler := handler.NewNotFoundRequestHandler()

	router := chi.NewRouter()
	router.Use(middleware.RealIP, middleware.Logger, middleware.Recoverer)
	router.Get(`/ping`, pingRequestHandler.HandleRequest)
	router.NotFound(notFoundHandler.HandleRequest)

	srv := &http.Server{Addr: ":8080", Handler: router}
	app := NewApp(srv)

	return &Container{
		app: app,
	}
}
