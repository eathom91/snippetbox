package main

// Main entry point for application execution.
import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application struct will manage project's dependency needs
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Creating, parsing flags from CLI
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()
	// Creating info, error loggers
	infoLog := log.New(os.Stdin, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdin, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// Creating application obj with logger intialization.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	// Creating the server with the port address, errorLogger, and mux with all our routes.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	// Info logging, and begin server listening.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
