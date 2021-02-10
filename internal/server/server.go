package server

import (
	"github.com/apalyukha/image-generator/configs"
	"github.com/apalyukha/image-generator/pkg/img"
	"log"
	"net/http"
	"strconv"
)

func Run(conf configs.IConf) {
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
		log.Println(err)
	}
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := img.GenerateFavicon()
	if err != nil {
		log.Println(err)
	}
	//rend(w, "favicon")

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err = w.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
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
