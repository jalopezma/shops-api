package main

import (
	"log"
	"os"

	"github.com/jalopezma/shops-api/config"
	"github.com/jalopezma/shops-api/handlers"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Port: %v\n", port)

	env := os.Getenv("APP_ENV")
	if env == "PROD" {
		log.Printf("APP_ENV: Production %q", env)
	} else if env == "DEV" {
		log.Printf("APP_ENV: Development %q", env)
	} else {
		env = "dev"
	}
	config.InitDeps(env)

	router := handlers.NewRouter()

	n := negroni.Classic()
	// CORS for * or https://*.foo.com origins, allowing:
	options := cors.New(cors.Options{
		//		AllowedOrigins:   []string{"http://localhost:8000", "http://localhost:4200", "*.stratabet.com", "*.stratasport.com", "http://stratasport.com", "*stratatips.co", "*stratapro.co"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Encoding", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token", "Origin", "Upgrade", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	})

	n.Use(options)
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(router)
	n.Run(":" + port)
}
