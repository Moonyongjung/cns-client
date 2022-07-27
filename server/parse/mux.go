package parse

import (
	"html/template"
	"net/http"
)

func ClientRequestMux(mux *http.ServeMux, templates *template.Template) *http.ServeMux{
	mux = static(mux, templates) 
	mux = wallet(mux, templates)
	mux = domain(mux, templates)
	mux = retrieve(mux, templates)
	mux = send(mux, templates)

	return mux
}