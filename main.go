package main

import (
	"context"
	"fmt"
	"gochicoba/db"
	"gochicoba/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database := db.DatabaseInitialize()
	addr := os.Getenv("APP_PORT")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: InitializeRoute(database),
	}

	go func() {
		server.ListenAndServe()
	}()

	defer Stop(server)
	log.Printf("Started server on : %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch), "in server")
	log.Println("Stopping API Server")
}

func InitializeRoute(db *gorm.DB) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.MethodNotAllowed(handler.MethodNotAllowedHandler)
	router.NotFound(handler.NotFoundHandler)

	uh := BuyHandler(db)

	router.Route("/buy", func(router chi.Router) {
		router.Get("/", uh.GetAllBuys)
		// router.Post("/", uh.CreateBuy)
		router.Route("/transactionx", func(r chi.Router) {
			r.Post("/", uh.CreateTransaction)
		})
	})

	return router
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server corectly: %v\n", err)
		os.Exit(1)
	}
}
