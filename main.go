package main

import (
	"net/http"
	"html/template"
)

type page struct {
	Title string
	Msg string
}

func hello(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html")
    t, _ := template.ParseFiles("tmpl/index.html")
    t.Execute(w, &page{Title: "Моментальные переводы денег онлайн между картами любых банков Казахстана Card2Card"})	
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":9000", nil)
}