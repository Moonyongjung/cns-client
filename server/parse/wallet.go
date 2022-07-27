package parse

import (
	"html/template"
	"net/http"

	"github.com/Moonyongjung/cns-client/rest"
	cns "github.com/Moonyongjung/cns-client/types"
	"github.com/Moonyongjung/cns-client/util"	
	"github.com/mitchellh/mapstructure"
)

func wallet(mux *http.ServeMux, templates *template.Template) *http.ServeMux{
	gatewayEndpoint := util.GetConfig().Get("gatewayEndpoint")
	mux.HandleFunc("/wallet/login/", func(w http.ResponseWriter, r *http.Request) {									
		err := templates.ExecuteTemplate(w, "inputmne.html", nil)
		if err != nil {
			util.LogGw(err)
		}
	})
	mux.HandleFunc("/wallet/create/", func(w http.ResponseWriter, r *http.Request) {	
		targetUrl := gatewayEndpoint+"/client/wallet/create/"
		util.LogGw("Target URL : ", targetUrl)		
		responseBody, _ := rest.HttpClient("GET", targetUrl, nil, r)		

		var createMnemonicResponse cns.CreateMnemonicResponse
		createMnemonicResponseData := util.JsonUnmarshalData(createMnemonicResponse, responseBody)
		mapstructure.Decode(createMnemonicResponseData, &createMnemonicResponse)
		
		err := templates.ExecuteTemplate(w, "createmne.html", createMnemonicResponse)
		if err != nil {
			util.LogGw(err)
		}
	})
	mux.HandleFunc("/wallet/address/", func(w http.ResponseWriter, r *http.Request) {									
		targetUrl := gatewayEndpoint+"/client/wallet/address/"
		util.LogGw("Target URL : ", targetUrl)
		util.LogGw("Request data : ", r.FormValue("mnemonic"))		
		responseBody, cookie := rest.HttpClient("POST", targetUrl, []byte(r.FormValue("mnemonic")), r)
		util.LogGw("server response : ", string(responseBody))		

		SessionManage(w, cookie)		

		var addressResponse cns.AddressResponse
		addressResponseData := util.JsonUnmarshalData(addressResponse, responseBody)
		mapstructure.Decode(addressResponseData, &addressResponse)
		
		err := templates.ExecuteTemplate(w, "address.html", addressResponse)
		if err != nil {
			util.LogGw(err)
		}
	})

	return mux
}