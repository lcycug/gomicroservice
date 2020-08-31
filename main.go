package main

import (
	"github.com/joho/godotenv"
	"github.com/lcycug/gomicroservice/server"
	"github.com/lcycug/homepage"
	"log"
	"net/http"
	"os"
)

var (
	KeyFile     string
	ServiceAddr string
	CertFile    string
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	CertFile = os.Getenv("CERT_FILE")
	KeyFile = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("PORT")
}

func main() {
	logger := log.New(os.Stdout, "DEBUG | ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandler(logger)
	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	logger.Println("server starting...")
	srv := server.New(mux, ServiceAddr)

	err := srv.ListenAndServeTLS(CertFile, KeyFile)

	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
