package rest

import (
	"bytes"
	"crypto/tls"	
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Moonyongjung/cns-client/util"	
)

//-- Get account number and sequence
func HttpClient(method string, url string, body []byte, r *http.Request) ([]byte, *http.Cookie){
	var request *http.Request
	var err error

	if method == "GET" {
		request, err = http.NewRequest("GET", url, nil)
		if err != nil {
			util.LogHttpClient(err)
		}
	} else {
		buf := bytes.NewBuffer(body)
		request, err = http.NewRequest("POST", url, buf)
		if err != nil {
			util.LogHttpClient(err)
		}
	}	

	cookies := r.Cookies()	

	if len(cookies) != 0 {
		for _, cookie := range cookies {
			if cookie.Name == "session" {
				request.AddCookie(cookie)
			}
		}		
	}	

	hClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	hClient.Timeout = time.Second * 30
	defer func() {
		if err := recover(); err != nil {
			util.LogHttpClient(err)		
		}
	}()
	response, err := hClient.Do(request)
	if err != nil {
		util.LogHttpClient(err)
	}

	cookies = response.Cookies()
	var cnsCookie *http.Cookie

	if len(cookies) != 0 {
		for _, cookie := range cookies {
			if cookie.Name == "session" {
				cnsCookie = cookie
			}
		}		
	}
	
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		util.LogHttpClient(err)
	}

	response.Body.Close()

	if len(cookies) != 0 {
		return responseBody, cnsCookie
	}
	return responseBody, nil
}