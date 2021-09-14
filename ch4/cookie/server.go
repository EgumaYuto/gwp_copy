package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "first cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "second cookie",
		Value:    "Manning Publications Go",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cs := r.Cookies()
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
