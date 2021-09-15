package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := selectTemplate()
	t.ExecuteTemplate(w, "layout", "")
}

func selectTemplate() (*template.Template, error) {
	rand.Seed(time.Now().Unix())
	if rand.Intn(10) > 5 {
		return template.ParseFiles("layout.html", "hello_red.html")
	} else {
		return template.ParseFiles("layout.html")
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
