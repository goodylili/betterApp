package server

import (
	"BetterApp/internal/http"
	"BetterApp/internal/models"
	"BetterApp/internal/users"
	"fmt"
	"log"
)

// Run - is going to be responsible for / the instantiation and startup of our / go application
func Run() error {
	fmt.Println("starting up the application...")

	store, err := models.NewDatabase()
	if err != nil {
		log.Println("Database Connection Failure")
		return err
	}

	if err := store.MigrateDB(); err != nil {
		log.Println("failed to setup store migrations")
		return err
	}

	userService := users.NewService(store)

	handler := http.NewHandler(userService)

	if err := handler.Serve(); err != nil {
		log.Println("failed to gracefully serve our application")
		return err
	}

	return nil

}
func main() {
	fmt.Println("Betterstack Go REST API Tutorial")
	if err := Run(); err != nil {
		log.Println(err)
	}

}
