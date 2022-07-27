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

func send(mux *http.ServeMux, templates *template.Template) *http.ServeMux{
	gatewayEndpoint := util.GetConfig().Get("gatewayEndpoint")
	clientPort := util.GetConfig().Get("clientPort")
	mux.HandleFunc("/send/index/", func(w http.ResponseWriter, r *http.Request) {									
		targetUrl := gatewayEndpoint+"/client/send/index/"
		util.LogGw("Target URL : ", targetUrl)		
		
		for _, value := range r.Cookies() {
			util.LogGw(value.Name)
			util.LogGw(value.Value)
		}
		responseBody, _ := rest.HttpClient("GET", targetUrl, nil, r)
		
		util.LogGw(string(responseBody))				

		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)
		
		err := templates.ExecuteTemplate(w, "sendindex.html", httpResponseStruct)
		if err != nil {
			util.LogGw(err)
		}
	})
	mux.HandleFunc("/send/inquiry/", func(w http.ResponseWriter, r *http.Request) {					

		targetDomain := r.FormValue("domain")
		amount := r.FormValue("amount")

		requestData := targetDomain + "," + amount

		targetUrl := gatewayEndpoint+"/client/send/inquiry/"
		util.LogGw("Target URL : ", targetUrl)
		util.LogGw("Request data : ", requestData)		
		responseBody, cookie := rest.HttpClient("POST", targetUrl, []byte(requestData), r)
		util.LogGw(string(responseBody))

		SessionManage(w, cookie)

		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)		

		pageConvertUrl := "http://localhost:"+clientPort+"/send/confirmed"+
			"?resCode="+util.ToString(httpResponseStruct.ResCode, "")+
			"&resMsg="+url.QueryEscape(util.ToString(httpResponseStruct.ResMsg, ""))
		util.LogGw(pageConvertUrl)

		http.Redirect(w, r, pageConvertUrl, http.StatusFound)
	})
	mux.HandleFunc("/send/confirmed/", func(w http.ResponseWriter, r *http.Request) {	
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
		

		err = templates.ExecuteTemplate(w, "sendconfirm.html", httpResponseStruct)
		if err != nil {
			util.LogGw(err)
		}
	})

	return mux
}