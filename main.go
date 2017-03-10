package main

import (
	"log"
	"net/http"
)

const (
	WallPaperPath string = "./wallpapers/"
	WallPaperUrl  string = "/wallpapers/"
	ListenPort    string = ":1234"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(WallPaperPath))))

	log.Printf("WallPaper Server Ready Start\n")
	log.Printf("Use 127.0.0.1%s To Visit\n", ListenPort)

	if err := http.ListenAndServe(ListenPort, nil); err != nil {
		log.Fatalln("WallPaper Server Start Error: ", err)
	}
}
