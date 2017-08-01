package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestEcho generates a random 8KB POST request and confirms that
// echo returns the text that is passed to it.
func TestEcho(t *testing.T) {
	buf := make([]byte, 8192)
	rand.Read(buf)
	body := bytes.NewBuffer(buf)

	req, err := http.NewRequest("POST", "/echo", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	echo(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("echo returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := string(buf)
	if body := rr.Body.String(); body != expected {
		t.Errorf("echo returned unexpected body: got %v bytes want %v bytes", len(body), len(expected))
	}
}
