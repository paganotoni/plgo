package handlers

import (
	_ "embed"
	"net/http"

	"github.com/gobuffalo/plush/v4"
)

//go:embed landing.html
var landingHTML string

func Landing(w http.ResponseWriter, r *http.Request) {
	c := plush.NewContextWithContext(r.Context())
	c.Set("title", "This is a plush template")
	c.Set("items", []string{"one", "two", "three"})

	s, err := plush.Render(landingHTML, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte(s))
}
