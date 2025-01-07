package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dzemildupljak/go_app_tools/internal/application"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Login endpoint")
}

func registerHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(w, "Register endpoint")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := application.GetUsers(r.Context())

	if err != nil {
		http.Error(w, "Bad request: error in user details", http.StatusBadRequest)
		return
	}
	//utils.LogError(r.Context(),
	//	fmt.Errorf("error in user details"), "Get user details error")

	// Return a 400 Bad Request with an error message
	//http.Error(w, "Bad request: error in user details", http.StatusBadRequest)

	// Return a 200 OK response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// Encode the users slice into JSON and write it to the response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func registerRoutes(router *mux.Router) {
	// Auth routes
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	authRouter.HandleFunc("/register", registerHandler).Methods(http.MethodPost)

	// User routes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(httpLoggingMiddleware)
	userRouter.HandleFunc("/{id}", getUserHandler).Methods(http.MethodGet)
}

func NewServer() *http.Server {
	router := mux.NewRouter()

	registerRoutes(router)

	// Configure server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server will run at http://localhost%s\n", server.Addr)
	return server
}
