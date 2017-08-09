package main

import (
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"

	_ "github.com/wedow/homework/db"
)

var domain string = "test.wedow.cc"

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
}

func main() {
	http.Handle("/", http.FileServer(assetFS()))
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/text", textHandler)

	log.Println("Listening on :8080")
	go http.ListenAndServe(":8080", nil)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain), //your domain here
		Cache:      autocert.DirCache("certs"),     //folder for storing certificates
	}

	server := &http.Server{
		Addr: ":8443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	log.Println("Listening on :8443", domain)
	log.Fatal(server.ListenAndServeTLS("", "")) //key and cert are comming from Let's Encrypt
}
