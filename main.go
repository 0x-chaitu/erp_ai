package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	firebase "firebase.google.com/go"
	"github.com/0x-chaitu/rag_erp/internal/api"
	"github.com/0x-chaitu/rag_erp/pkg/utils"
	"google.golang.org/api/option"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := utils.NewDatabasePool(ctx, 30)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	opt := option.WithCredentialsFile("rag-erp-firebase.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}

	api := api.NewAPI(ctx, db, authClient)
	srv := api.Server(8000)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Println("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	cancel()
	srv.Shutdown(ctx)

}
