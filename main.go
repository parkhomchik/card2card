package main

import (
	"net/http"
	"html/template"
	"fmt"
	"strconv"
)

type page struct {
	Title string
	Msg string
}

func index(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html")
    t, _ := template.ParseFiles("tmpl/index.html")
    t.Execute(w, &page{Title: "Моментальные переводы денег онлайн между картами любых банков Казахстана Card2Card"})	
}

func p2p(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("tmpl/index.html")
        t.Execute(w, &page{Title: "Моментальные переводы денег онлайн между картами любых банков Казахстана Card2Card"})
    } else {
        r.ParseForm()
		senderCard := ""
		for i := 1; i <=4; i++ {
			sc := r.Form["sender_card_num_"+strconv.Itoa(i)]
			senderCard += sc[0]
		}

		//senderCard := r.Form["sender_card_num_1"]
		//senderCard = append(senderCard, r.Form["sender_card_num_2"])

		/*senderCard1 := r.Form["sender_card_num_1"]
		senderCard2 := r.Form["sender_card_num_2"]
		senderCard3 := r.Form["sender_card_num_3"]
		senderCard4 := r.Form["sender_card_num_4"]

		senderCard := senderCard1[0] + senderCard2[0] + senderCard3[0] + senderCard4[0]*/


        fmt.Println("sender card:", senderCard)
    }
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
    http.HandleFunc("/index", index)
	http.HandleFunc("/", index)
	http.HandleFunc("/p2p", p2p)
    http.ListenAndServe(":9000", nil)
}