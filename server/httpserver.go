package server

import (
	"html/template"
	"net/http"
	
	"github.com/Moonyongjung/cns-client/util"
	"github.com/Moonyongjung/cns-client/server/parse"
	
	"github.com/rs/cors"
)

//-- HTTPServer operates for sending or invoking transaction when user call
func RunHttpServer() {
	clientPort := util.GetConfig().Get("clientPort")
	templates := template.Must(template.ParseFiles(
		"./templates/index.html", 
		"./templates/signup.html",
		"./templates/inputmne.html",
		"./templates/createmne.html",
		"./templates/address.html",
		"./templates/mapping.html",
		"./templates/confirm.html",
		"./templates/retrieveindex.html",
		"./templates/retrieveinfo.html",
		"./templates/sendindex.html",
		"./templates/sendconfirm.html",
	))
	mux := http.NewServeMux()	
	mux = parse.ClientRequestMux(mux, templates)	
	handler := cors.Default().Handler(mux)
	util.LogHttpServer("Server Listen...")	

	err := http.ListenAndServe(":"+clientPort, handler)
	if err != nil {
		util.LogHttpServer(err)	
	}
}