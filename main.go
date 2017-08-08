package main

import (
	"log"
	"net/http"

	_ "github.com/wedow/homework/db"
)

var port string = ":8080"

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
}

func main() {
	http.Handle("/", http.FileServer(assetFS()))
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/text", textHandler)

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
