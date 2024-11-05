package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found")
	} 
	fmt.Println(`Port:`, port)

	router:= chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // You can specify allowed origins here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router:= chi.NewRouter()

	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)

	srv:= &http.Server{
		Handler: router,
		Addr: ":"+port,
	}

	log.Printf("Server running at Port: %v", port)
	err:= srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	
}