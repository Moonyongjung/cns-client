package parse

import (
	"net/http"

	"github.com/Moonyongjung/cns-client/util"
)

func SessionManage(w http.ResponseWriter, cookie *http.Cookie) {
	util.LogGw("cookie string : ", cookie.String())	
	http.SetCookie(w, cookie)	
}