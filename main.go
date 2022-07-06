package main

import (
	"html/template"
	"log"
	"net/http"
)

/*

 */
func main() {
	server := http.Server{Addr: ":8080"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", idx)
	err := server.ListenAndServe()
	log.Fatalln(err)
}

func idx(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(res, nil)
}
