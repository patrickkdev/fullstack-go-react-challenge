package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"api/configs"
	"api/internal/application"
	"api/internal/communication/api"
	"api/internal/infrastructure/db"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	dbConn, err := db.Connect(configs.DBConfig)
	if err != nil {
		panic(err)
	} else {
		log.Println("connected to database")
	}

	userRepo := db.NewUserRepository(dbConn)
	jobRepo := db.NewJobRepository(dbConn)
	jobAppRepo := db.NewJobApplicationRepository(dbConn)

	authService := application.NewAuthService(userRepo)
	userService := application.NewUserService(userRepo)
	jobService := application.NewJobService(jobRepo)
	jobAppService := application.NewJobApplicationService(jobAppRepo)

	r := api.NewRouter(authService, userService, jobService, jobAppService)

	server := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server: %v", err)
		}
	}()

	log.Println("server started on :4000")

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exiting")
}
