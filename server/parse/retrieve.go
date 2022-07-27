package parse

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/Moonyongjung/cns-client/rest"
	cns "github.com/Moonyongjung/cns-client/types"
	"github.com/Moonyongjung/cns-client/util"	
	"github.com/mitchellh/mapstructure"
)

func retrieve(mux *http.ServeMux, templates *template.Template) *http.ServeMux{
	gatewayEndpoint := util.GetConfig().Get("gatewayEndpoint")
	mux.HandleFunc("/retrieve/index/", func(w http.ResponseWriter, r *http.Request) {	
		err := templates.ExecuteTemplate(w, "retrieveindex.html", nil)
		if err != nil {
			util.LogGw(err)
		}
	})
	mux.HandleFunc("/retrieve/info/", func(w http.ResponseWriter, r *http.Request) {	
		var domainMappingRequest cns.DomainMappingRequest
		var targetUrl string

		util.LogGw(r.FormValue("domain"))
		

		if len(r.FormValue("domain")) <= 20 {
			domainMappingRequest = cns.DomainMappingRequest{
				DomainName: r.FormValue("domain"),
			}
			targetUrl = gatewayEndpoint+"/api/wasm/cns-query-by-domain"
		} else {
			domainMappingRequest = cns.DomainMappingRequest{
				AccountAddress: r.FormValue("domain"),
			}
			targetUrl = gatewayEndpoint+"/api/wasm/cns-query-by-account"
		}
		
		requestData, err := util.JsonMarshalData(domainMappingRequest)
		if err != nil {
			util.LogGw(err)
		}

		util.LogGw("Target URL : ", targetUrl)
		util.LogGw("Request data : ", string(requestData))		

		responseBody, _ := rest.HttpClient("POST", targetUrl, requestData, r)
		util.LogGw(string(responseBody))
		
		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)	

		if httpResponseStruct.ResCode == 0 {
			splitData := strings.Split(httpResponseStruct.ResData, ":")
			convertStr := strings.Replace(splitData[2], "\"", "", -1)
			convertStr = strings.Replace(convertStr, "}", "", -1)

			httpResponseStruct.ResData = convertStr
			util.LogGw(convertStr)	

			err := templates.ExecuteTemplate(w, "retrieveinfo.html", httpResponseStruct)
			if err != nil {
				util.LogGw(err)
			}
		} else {
			err := templates.ExecuteTemplate(w, "retrieveinfo.html", httpResponseStruct)
			if err != nil {
				util.LogGw(err)
			}
		}
	})

	return mux
}