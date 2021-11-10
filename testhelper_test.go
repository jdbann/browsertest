package browsertest

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var templates *template.Template

func init() {
	var err error

	templates, err = template.ParseGlob("testdata/*.tmpl.html")
	if err != nil {
		log.Fatal(err)
	}
}

func newTestServer(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		value := r.PostFormValue("value")

		w.WriteHeader(http.StatusOK)
		templates.ExecuteTemplate(w, "actions.tmpl.html", value)
	})

	mux.HandleFunc("/queries", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		templates.ExecuteTemplate(w, "queries.tmpl.html", nil)
	})

	server := httptest.NewServer(mux)

	t.Cleanup(func() { server.Close() })

	return server
}
