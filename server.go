package main

import "net/http"

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://robotikazabulgaria.com", http.StatusFound)
}
