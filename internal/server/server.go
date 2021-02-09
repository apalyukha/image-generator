package server

import (
	"github.com/apalyukha/image-generator/configs"
	"log"
	"net/http"
)

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	log.Println("Server starting ...")
	if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
		log.Fatalln(err)
	}

}

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatalln(err)
	}
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon")
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "img")

	// Щоб не було дублювання коду, виносимо провірку у func rend()
	//------------------------------------------------------------

	//_, err := w.Write([]byte("img"))
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
