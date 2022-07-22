package main

// Routes.go contains all the logic for registering routes in a mux that is then returned.
// Routes.go makes use of the handlers.go functions to initialize each route
import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
