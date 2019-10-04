// Author : Teknas

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",Hello)
	http.ListenAndServe(":8080",nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World !")
}
