package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	conn, err := postgres.New(postgres.Settings{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "15370933",
		DBName:   "clean",
	})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	fmt.Println("Connected to PostgreSQL database successfully")

	contactRepo := internal.NewRepository(conn.Pool)
	contactUseCase := internal.NewUseCase(contactRepo)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	contactDelivery := internal.NewDelivery(contactUseCase, logger)

	http.Handle("/contacts", contactDelivery)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("HTTP server error: ", err)
		}
	}()

	fmt.Println("HTTP server started on port 8080")

	select {}
}
