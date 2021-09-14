package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	html := `<html>
		<head>
			<title>Go Web Progrramming</title>
		</head>
		<body>
			<h1>Hello World</h1>
		</body>
	</html>`
	w.Write([]byte(html))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。ほかを当たってください。")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func jsonExeample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Cabos",
		Threads: []string{"1番目", "2番目", "3番目"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExeample)
	server.ListenAndServe()
}
