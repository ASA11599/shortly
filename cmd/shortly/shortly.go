package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/ASA11599/shortly/internal/handlers"
	"github.com/ASA11599/shortly/internal/storage"
	"github.com/go-chi/chi/v5"
)

func interruptChannel() chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	return signals
}

func getHost() string {
	if host, ok := os.LookupEnv("HOST"); ok {
		return host
	}
	return "0.0.0.0"
}

func getPort() int {
	if port, ok := os.LookupEnv("PORT"); ok {
		p, err := strconv.ParseInt(port, 10, 16)
		if err == nil {
			return int(p)
		}
	}
	return 80
}

func main() {

	mux := chi.NewRouter()

	var store storage.Store = storage.NewMemoryStore()
	defer store.Close()

	mux.Get("/{alias}", handlers.NewGetHandler(store).ServeHTTP)
	mux.Post("/", handlers.NewPostHandler(store).ServeHTTP)

	go panic(http.ListenAndServe(fmt.Sprintf("%s:%d", getHost(), getPort()), mux))

	<-interruptChannel()

}
