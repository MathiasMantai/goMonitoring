package main

/*
	project: goMonitoring
	version: 1.1
	author: Mathias Mantai <mmantaibusiness@gmail.com>
*/

import (
	"encoding/json"
	"fmt"
	"github.com/mathiasmantai/goMonitoring/src"
	"html/template"
	"log"
	"net/http"
)

const (
	port = "8080"
)

type PageData struct {
	Title        string
	Body         interface{}
	CurrentRoute string
}

type CpuData struct {
	CpuUsage      float64
	VirtualMemory float64
}

func serveStaticFiles() {
	//serve css
	styles := http.FileServer(http.Dir("./web/css"))
	http.Handle("/css/", http.StripPrefix("/css/", styles))

	//serve javascript
	js := http.FileServer(http.Dir("./web/js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	//serve images
	img := http.FileServer(http.Dir("./web/img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	pages := http.FileServer(http.Dir("./web/pages"))
	http.Handle("/pages/", http.StripPrefix("/pages/", pages))
}

func main() {

	pages := []string{
		"web/index.html",
		"web/pages/cpu.html",
		"web/pages/container.html",
		"web/pages/network.html",
	}

	//default route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:        "Monitoring",
			Body:         src.HostInfo(),
			CurrentRoute: "/",
		}
		renderTemplateWithContent(w, data, pages...)
	})

	http.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		data := CpuData{CpuUsage: src.CPUData(false), VirtualMemory: src.VirtualMemory()}
		response, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(response)
	})

	//container route
	http.HandleFunc("/container", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:        "Monitoring",
			Body:         src.ContainerData(),
			CurrentRoute: "/container",
		}
		renderTemplateWithContent(w, data, pages...)
	})

	//network route
	http.HandleFunc("/network", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:        "Monitoring",
			Body:         "Test",
			CurrentRoute: "/network",
		}
		renderTemplateWithContent(w, data, pages...)
	})

	//serve static files
	serveStaticFiles()

	fmt.Println("starting webserver on port " + port)
	http.ListenAndServe(":"+port, nil)
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

func renderTemplateWithContent(w http.ResponseWriter, data interface{}, content ...string) {
	t, err := template.ParseFiles(content...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
