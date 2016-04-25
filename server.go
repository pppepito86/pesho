package main

import "net/http"
import "os"
import "path/filepath"

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/lom.zip" {
		w.Header().Set("Content-Disposition", "attachment; filename=lom.zip")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		pwd, _ := os.Getwd()
		file := pwd + string(filepath.Separator) + "lom.zip"
		http.ServeFile(w, r, file)
	} else {
		http.Redirect(w, r, "http://robotikazabulgaria.com", http.StatusFound)
	}
}
