package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func configureRoutes() *chi.Mux {
	router := chi.NewRouter()

	/**
	 * Configure router.
	 *
	 * RedirectSlashes - Redirect slashes on routing paths.
	 * DefaultCompress - Gzip compression for clients that accept compressed responses.
	 * Recover - Gracefully absorb panics and prints the stack trace.
	 * Logger - Logs the start and end of each request with the elapsed processing time.
	 */
	router.Use(
		middleware.RedirectSlashes,
		middleware.DefaultCompress,
		middleware.Recoverer,
		middleware.Logger,
	)

	// Mount resources here.

	return router
}

func main() {
	// Push Routes() inside router here.
}
