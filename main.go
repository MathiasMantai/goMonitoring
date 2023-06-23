package main

import (
	"html/template"
	"net/http"
	"fmt"
)

const (
	port = "8080"
)

type PageData struct {
	Title string
	Body string
}

func main() {	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData {
			Title: "Monitoring",
			Body: "Test",
		}
		renderTemplate(w, "web/index.html", data)
	})

	//serve css
	styles := http.FileServer(http.Dir("./web/css"))
	http.Handle("/css/", http.StripPrefix("/css/", styles))

	//serve javascript
	js := http.FileServer(http.Dir("./web/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	//serve images
	img := http.FileServer(http.Dir("./web/img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	//serve external libraries
	img := http.FileServer(http.Dir("./web/img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	fmt.Println("starting webserver on port " + port)
	http.ListenAndServe(":" + port, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}