package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

}
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
