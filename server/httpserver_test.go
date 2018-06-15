package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/stat", nil)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: except: %v, got %v\n", http.StatusOK, status)
	}

	expected := `{"current_connection_count":0,"current_request_rate":0,"processed_request_count":0,"current_request_count":0,"remaining_jobs":0,"current_goroutine_count":2,"current_connected_client":null}`

	if rr.Body.String() != expected {
		t.Errorf("unexpected body: expect: %v, got %v\n", expected, rr.Body.String())
	}
}
