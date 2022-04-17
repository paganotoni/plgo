package server

import (
	"fmt"
	"net/http"

	"github.com/paganotoni/plgo/web/landing"
)

func Start() error {
	http.HandleFunc("/", landing.HandlerFn)
	fmt.Println(">>> Application running on http://localhost:8080 <<<")

	return http.ListenAndServe(":8080", nil)
}
