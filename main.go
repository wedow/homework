package main

import (
	"io"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
}

func main() {
	http.Handle("/", http.FileServer(assetFS()))
	http.HandleFunc("/echo", echo)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	size, err := io.Copy(rw, req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("[INFO] wrote %d bytes", size)
}
