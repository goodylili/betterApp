package http

import (
	"BetterApp/internal/users"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Handler struct {
	Router *mux.Router
	Users  users.UserService
	Server *http.Server
}

// NewHandler - returns a pointer to a Handler
func NewHandler(users users.UserService) *Handler {
	log.Println("setting up our handler")
	h := &Handler{
		Users: users,
	}

	h.Router = mux.NewRouter()

	h.mapRoutes()

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080", // Good practice to set timeouts to avoid Slow-loris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/user/create", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/api/v1/users/{id}", h.GetUserByID).Methods("GET")
	h.Router.HandleFunc("/api/v1/users/{id}", h.UpdateUser).Methods("PUT")
	h.Router.HandleFunc("/api/v1/users/{id}", h.DeleteUser).Methods("DELETE")
}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	// CreateAccount a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}
