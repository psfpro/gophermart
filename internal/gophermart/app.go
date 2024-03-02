package gophermart

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp(httpServer *http.Server) *App {
	return &App{httpServer: httpServer}
}

func (a *App) Run() {
	a.runHTTPServer()
	a.waitSignal()
}

func (a *App) runHTTPServer() {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()
}

func (a *App) waitSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-signalChan
	signal.Stop(signalChan)
	log.Printf("received signal %s, shutting down", sig.String())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.shutdown(ctx)
}

func (a *App) shutdown(ctx context.Context) {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("shutdown http server error %v", err)
	}
}
