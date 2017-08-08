package main

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/wedow/homework/db"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes generates a random string of length n for tests
func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

// TestEcho generates a random 8KB POST request and confirms that
// echo returns the text that is passed to it.
func TestEchoHandler(t *testing.T) {
	buf := RandStringBytes(8192)
	body := bytes.NewBuffer(buf)

	req, err := http.NewRequest("POST", "/echo", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	echoHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("echo returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := string(buf)
	if body := rr.Body.String(); body != expected {
		t.Errorf("echo returned unexpected body: got %v bytes want %v bytes", len(body), len(expected))
	}
}

func TestTextHandler(t *testing.T) {
	// generate random request body
	buf := RandStringBytes(8192)
	text, _ := json.Marshal(map[string]interface{}{"Content": string(buf)})
	body := bytes.NewBuffer(text)

	// create request
	req, err := http.NewRequest("POST", "/echo", body)
	if err != nil {
		t.Fatal(err)
	}

	// open mock database
	database, mock, err := sqlmock.New()
	db.Use(database)
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt"}).AddRow(1, time.Now())
	mock.ExpectQuery("INSERT INTO posts").WithArgs(string(buf)).WillReturnRows(rows)

	// run request handler
	rr := httptest.NewRecorder()
	textHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("text returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var result map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &result)

	if _, ok := result["Id"]; !ok {
		t.Errorf("text did not populate Id field")
	}

	if _, ok := result["CreatedAt"]; !ok {
		t.Errorf("text did not populate CreatedAt field")
	}

	expected := string(buf)
	if body, ok := result["Content"]; !ok {
		t.Errorf("text did not populate Content field")
	} else if body := body.(string); body != expected {
		t.Errorf("text returned unexpected body: got %v bytes want %v bytes", len(body), len(expected))
	}
}
