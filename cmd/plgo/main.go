package main

import (
	"fmt"
	"net/http"

	"github.com/paganotoni/plgo/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Landing)
	fmt.Println(">>> Application running on http://localhost:8080 <<<")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
