package handlers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/ASA11599/shortly/internal/core"
	"github.com/ASA11599/shortly/internal/storage"
	"github.com/go-chi/chi/v5"
)

type GetHandler struct {
	s storage.Store
}

func NewGetHandler(s storage.Store) http.Handler {
	return &GetHandler{ s: s }
}

func (gh *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	link := core.Expand(gh.s, alias)
	if link == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found."))
		return
	}
	http.Redirect(w, r, link, http.StatusFound)
}

type PostHandler struct {
	s storage.Store
}

func NewPostHandler(s storage.Store) http.Handler {
	return &PostHandler{ s: s }
}

func (ph *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	linkBytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error reading link."))
			return
		}
		link := string(linkBytes)
		_, err = url.ParseRequestURI(link)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid link."))
			return
		}
		alias := core.Shorten(ph.s, link)
		w.Write([]byte(alias))
}
