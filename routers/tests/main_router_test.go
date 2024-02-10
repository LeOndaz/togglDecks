package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"togglDecks/routers"
)

func TestPingRouter(t *testing.T) {
	router := routers.SetupMainRouter()

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/ping", nil)

	if err != nil {
		t.Fatal(err.Error())
	}

	router.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Fatal("/ping didn't respond with 200")
	}

	if recorder.Body.String() != "PONG" {
		t.Errorf("expected body %s; got %s", "PONG", recorder.Body.String())
	}
}
