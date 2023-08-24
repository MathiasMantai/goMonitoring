package main

/*
	project: goMonitoring
	version: 1.1
	author : Mathias Mantai <mmantaibusiness@gmail.com>
*/

import (
	"encoding/json"
	"fmt"
	"github.com/mathiasmantai/goMonitoring/src"
	"html/template"
	"log"
	"net/http"
	"time"
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
		renderTemplateWithContent(w, data, nil, pages...)
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

		//get url param
		urlParam := r.URL.Query()
		if urlParam["action"] != nil {
			action := urlParam["action"][0]
			containerId := urlParam["container"][0]

			if action == "start" {
				src.StartContainer(containerId)
			} else {
				src.StopContainer(containerId)
			}

			time.Sleep(time.Second)
		}

		//get pagedata
		data := PageData{
			Title:        "Monitoring",
			Body:         src.ContainerData(),
			CurrentRoute: "/container",
		}
		renderTemplateWithContent(w, data, nil, pages...)
	})

	//network route
	http.HandleFunc("/network", func(w http.ResponseWriter, r *http.Request) {
		inter := src.GetNetworkInterfaces()
		activeInterface := src.FilterNetworkInterfaces(&inter)
		data := PageData{
			Title:        "Monitoring",
			Body: activeInterface,
			CurrentRoute: "/network",
		}
		renderTemplateWithContent(w, data, nil, pages...)
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

func renderTemplateWithContent(w http.ResponseWriter, data interface{}, function func(interface{}) interface{}, content ...string) {
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
