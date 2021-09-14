package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, []string{"月", "火", "...", "日"})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
