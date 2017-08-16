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
	user := "foobar"

	// generate random request body
	buf := RandStringBytes(8192)
	text, _ := json.Marshal(map[string]interface{}{"content": string(buf), "username": user})
	body := bytes.NewBuffer(text)

	// create request
	req, err := http.NewRequest("POST", "/text", body)
	if err != nil {
		t.Fatal(err)
	}

	// open mock database
	database, mock, err := sqlmock.New()
	db.Use(database)

	id, timestamp := 15, time.Now()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt"}).AddRow(id, timestamp)
	mock.ExpectQuery("INSERT INTO posts").WillReturnRows(rows)

	// run request handler
	rr := httptest.NewRecorder()
	textHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("text returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var result map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &result)

	if resultId, ok := result["id"]; !ok {
		t.Errorf("text did not populate id field")
	} else if resultId.(float64) != float64(id) {
		t.Errorf("text returned wrong id value: got %v, expected %v", resultId, id)
	}
	t.Log(result["created_at"])
	if resultTimestamp, ok := result["created_at"]; !ok {
		t.Errorf("text did not populate created_at field")
	} else if timestring, ok := resultTimestamp.(string); ok {
		ts, err := time.Parse(time.RFC3339Nano, timestring)
		if err != nil {
			t.Errorf("invalid created_at format - %s", err)
		} else if !ts.Equal(timestamp) {
			t.Errorf("text returned wrong created_at value: got %v, expected %v", ts, timestamp)
		}
	}

	expected := string(buf)
	if body, ok := result["content"]; !ok {
		t.Errorf("text did not populate content field")
	} else if body := body.(string); body != expected {
		t.Errorf("text returned unexpected body: got %v bytes want %v bytes", len(body), len(expected))
	}
}
