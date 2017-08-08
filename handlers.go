package main

import (
	"encoding/json"
	"fmt"
	"github.com/wedow/homework/db"
	"io"
	"log"
	"net/http"
)

func echoHandler(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	size, err := io.Copy(rw, req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("[INFO] wrote %d bytes", size)
}

func textHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		posts, err := db.GetPosts()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err.Error())
			return
		}

		json.NewEncoder(rw).Encode(posts)
	case "POST":
		defer req.Body.Close()

		var post db.Post
		err := json.NewDecoder(req.Body).Decode(&post)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(rw, err.Error())
			return
		}

		if err := post.Save(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err.Error())
			return
		}

		json.NewEncoder(rw).Encode(post)
	}
}
