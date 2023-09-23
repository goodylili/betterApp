package http

import (
	"BetterApp/internal/users"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"strconv"
)

// CreateUser decodes a User object from the HTTP request body and then tries to create a new user in the database using the CreateUser method of the UserService interface. If the user is successfully created, it encodes and sends the created user as a response.
func (h *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var u users.User
	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		http.Error(writer, "Failed to decode request body", http.StatusBadRequest)
		log.Println("Failed to decode request body:", err)
		return
	}

	err := h.Users.CreateUser(request.Context(), &u)
	if err != nil {
		http.Error(writer, "Failed to create user", http.StatusInternalServerError)
		log.Println("Failed to create user:", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(u); err != nil {
		log.Panicln("Failed to encode response:", err)
	}
}

// GetUserByID extracts the id from the URL parameters and then fetches the user with that id from the database using the GetUserByID method of the UserService interface. If the user is found, it encodes and sends the user as a response.
func (h *Handler) GetUserByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.Users.GetUserByID(request.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		log.Panicln(err)
	}
}

// UpdateUser updates a user by ID.
func (h *Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Decode the request body to get the updated user information
	var u users.User
	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update user
	err = h.Users.UpdateUser(request.Context(), u, uint(id))
	if err != nil {
		http.Error(writer, "Failed to update user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode and send response
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		log.Panicln(err)
	}
}

// DeleteUser deletes a user by ID.
func (h *Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete user
	err = h.Users.DeleteUser(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, "Failed to delete user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Send a success response
	writer.WriteHeader(http.StatusNoContent)
}
