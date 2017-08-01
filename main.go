package main

import (
	"io"
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
	http.HandleFunc("/echo", echo)

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func echo(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	size, err := io.Copy(rw, req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("[INFO] wrote %d bytes", size)
}
