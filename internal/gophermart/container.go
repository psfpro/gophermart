package gophermart

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/psfpro/gophermart/internal/gophermart/application"
	"github.com/psfpro/gophermart/internal/gophermart/infrastructure/accrual"
	"github.com/psfpro/gophermart/internal/gophermart/infrastructure/api/http/handler"
	"github.com/psfpro/gophermart/internal/gophermart/infrastructure/authentication"
	"github.com/psfpro/gophermart/internal/gophermart/infrastructure/storage/postgres"
)

type Container struct {
	router http.Handler
	app    *App
}

func (c *Container) Router() http.Handler {
	return c.router
}

func (c *Container) App() *App {
	return c.app
}

func NewContainer() *Container {
	config := NewConfig()
	ctx := context.Background()

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
	// Repositories
	userRepository := postgres.NewUserRepository(db)
	userRepository.CreateTable(ctx)
	transactionRepository := postgres.NewTransactionRepository(db)
	transactionRepository.CreateTable(ctx)

	// Services
	authenticationService := authentication.NewService()
	userService := application.NewUserService(userRepository, authenticationService)
	transactionService := application.NewTransactionService(transactionRepository)

	// HTTP handlers
	pingRequestHandler := handler.NewPingRequestHandler(db)
	notFoundHandler := handler.NewNotFoundRequestHandler()
	userLoginHandler := handler.NewUserLoginRequestHandler(userService)
	userRegisterHandler := handler.NewUserRegisterRequestHandler(userService)
	balanceHandler := handler.NewBalanceRequestHandler(transactionService, authenticationService)
	orderUploadHandler := handler.NewOrderUploadRequestHandler(transactionService, authenticationService)
	ordersHandler := handler.NewOrdersRequestHandler(transactionService, authenticationService)
	withdrawHandler := handler.NewWithdrawRequestHandler(transactionService, authenticationService)
	withdrawalsHandler := handler.NewWithdrawalsRequestHandler(transactionService, authenticationService)

	router := chi.NewRouter()
	router.Use(middleware.RealIP, middleware.Logger, middleware.Recoverer)
	router.Get(`/api/ping`, pingRequestHandler.HandleRequest)
	router.Post(`/api/user/login`, userLoginHandler.HandleRequest)
	router.Post(`/api/user/register`, userRegisterHandler.HandleRequest)
	router.Get(`/api/user/balance`, balanceHandler.HandleRequest)
	router.Post(`/api/user/orders`, orderUploadHandler.HandleRequest)
	router.Get(`/api/user/orders`, ordersHandler.HandleRequest)
	router.Post(`/api/user/balance/withdraw`, withdrawHandler.HandleRequest)
	router.Get(`/api/user/withdrawals`, withdrawalsHandler.HandleRequest)
	router.NotFound(notFoundHandler.HandleRequest)

	srv := &http.Server{Addr: ":8080", Handler: router}
	accrualClient := accrual.NewClient(config.accrualAddress, transactionRepository)
	app := NewApp(srv, accrualClient)

	return &Container{
		router: router,
		app:    app,
	}
}
