package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestInfo(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", nil)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Did not get 200 response from /info request. Code: %d", w.Code)
	}

}

func TestSubscribe(t *testing.T) {

	s := httptest.NewServer(setupRouter())
	defer s.Close()

	//replace http with ws
	url := fmt.Sprintf("ws%s/subscribe", s.URL[4:])

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("URL = %s\n", url)
		t.Fatalf("Failed to Subscribe to websocket: %s\n", err.Error())
	}
	defer ws.Close()

	//maybe some more testing?

}
