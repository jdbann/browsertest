package browsertest

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var actionsHTML *template.Template

func init() {
	var err error

	actionsHTML, err = template.ParseFiles("testdata/actions.html")
	if err != nil {
		log.Fatal(err)
	}
}

func newTestServer(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		actionsHTML.Execute(w, nil)
	})

	server := httptest.NewServer(mux)

	t.Cleanup(func() { server.Close() })

	return server
}
