package landing

import (
	_ "embed"
	"net/http"

	"github.com/paganotoni/plgo/web"
)

//go:embed landing.plush.html
var landingHTML string

func HandlerFn(w http.ResponseWriter, r *http.Request) {
	dat, err := web.Render(r.Context(), landingHTML, map[string]interface{}{
		"title": "Plush is very cool",
		"features": []string{
			"It has a nice syntax",
			"Allows for dynamic content",
			"Its cool",
		},
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte(dat))
}
