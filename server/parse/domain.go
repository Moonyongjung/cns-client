package parse

import (
	"html/template"
	"net/http"
	"net/url"
	"strconv"	

	"github.com/Moonyongjung/cns-client/rest"
	cns "github.com/Moonyongjung/cns-client/types"
	"github.com/Moonyongjung/cns-client/util"
	"github.com/mitchellh/mapstructure"
)

func domain(mux *http.ServeMux, templates *template.Template) *http.ServeMux{
	gatewayEndpoint := util.GetConfig().Get("gatewayEndpoint")
	clientPort := util.GetConfig().Get("clientPort")
	mux.HandleFunc("/domain/mapping/", func(w http.ResponseWriter, r *http.Request) {	
		util.LogGw("requestURL : ", r.URL.RawQuery)
		userAddress := r.URL.RawQuery

		targetUrl := gatewayEndpoint+"/client/domain/mapping/"
		util.LogGw("Target URL : ", targetUrl)
		util.LogGw("Request data : ", userAddress)		
		responseBody, cookie := rest.HttpClient("POST", targetUrl, []byte(userAddress), r)
		util.LogGw("server response : ", string(responseBody))		

		SessionManage(w, cookie)

		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)	
		
		err := templates.ExecuteTemplate(w, "mapping.html", httpResponseStruct)
		if err != nil {
			util.LogGw(err)
		}
	})
	mux.HandleFunc("/domain/confirm/", func(w http.ResponseWriter, r *http.Request) {	
		util.LogGw("requestURL : ", r.URL.RawQuery)
		userAddress := r.URL.RawQuery
		requestData := userAddress + "," + r.FormValue("domain")

		targetUrl := gatewayEndpoint+"/client/domain/confirm/"
		util.LogGw("Target URL : ", targetUrl)
		util.LogGw("Request data : ", requestData)		
		responseBody, _ := rest.HttpClient("POST", targetUrl, []byte(requestData), r)

		util.LogGw("server response : ", string(responseBody))

		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)		

		pageConvertUrl := "http://localhost:"+clientPort+"/domain/confirmed"+
		"?resCode="+util.ToString(httpResponseStruct.ResCode, "")+
		"&resMsg="+url.QueryEscape(util.ToString(httpResponseStruct.ResMsg, ""))
		util.LogGw(pageConvertUrl)

		http.Redirect(w, r, pageConvertUrl, http.StatusFound)

	})
	mux.HandleFunc("/domain/confirmed/", func(w http.ResponseWriter, r *http.Request) {	
		util.LogGw("requestURL : ", r.URL.RawQuery)		

		urlParse, err := url.ParseQuery(r.URL.Query().Encode())
		if err != nil {
			util.LogGw(err)
		}
		resCode := urlParse.Get("resCode")
		resMsg := urlParse.Get("resMsg")

		var httpResponseStruct cns.HttpResponseStruct
		resCodeInt, err := strconv.Atoi(resCode)
		if err != nil {
			util.LogGw(err)
		}
		httpResponseStruct.ResCode = resCodeInt
		httpResponseStruct.ResMsg = resMsg		

		err = templates.ExecuteTemplate(w, "confirm.html", httpResponseStruct)
		if err != nil {
			util.LogGw(err)
		}
	})

	return mux
}