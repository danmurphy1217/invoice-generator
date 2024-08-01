package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danmurphy1217/invoice-generator/api"
	invoiceDB "github.com/danmurphy1217/invoice-generator/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
    database, _ := invoiceDB.Connect()

    err := database.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %q", err)
    }
    fmt.Println("Successfully connected to the database")

    r := chi.NewRouter()
    
    // Use some built-in middlewares
    r.Use(corsMiddleware)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Define API routes
    r.Get("/ping", api.HealthCheck)

    r.Post("/invoices", api.GenerateInvoiceHandler)
    
    // Start the server
    http.ListenAndServe(":8080", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}