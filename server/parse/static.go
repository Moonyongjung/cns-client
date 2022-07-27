package parse

import (
	"html/template"
	"net/http"

	"github.com/Moonyongjung/cns-client/util"
)

func static(mux *http.ServeMux, templates *template.Template) *http.ServeMux{
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {					
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {			
			util.LogGw(err)
		}
	})	

	return mux
}