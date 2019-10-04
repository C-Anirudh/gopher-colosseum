// Author: @C-Anirudh
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<h1>Hello, World</h1>`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8000", nil)
}
