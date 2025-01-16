package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test to verify the redirect from root ("/") to "/courses"
func TestRedirectRootToCourses(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Recorder to capture the response
	rr := httptest.NewRecorder()

	// Define the redirect handler to test
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/courses", http.StatusFound)
	})

	// Execute the handler
	handler.ServeHTTP(rr, req)

	// Verify the status code is HTTP 302 (Redirect)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	// Verify the Location header for the redirect
	location := rr.Header().Get("Location")
	if location != "/courses" {
		t.Errorf("handler returned wrong redirect location: got %v want %v", location, "/courses")
	}
}

