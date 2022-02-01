package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TwiN/go-color"
)

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, color.Ize(color.Green, "{\"status\": \"good\"}"))
	}

	req := httptest.NewRequest("GET", "http://local.services.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Fatal(color.Ize(color.Red, "Status Code Not OK"))
	}
}
