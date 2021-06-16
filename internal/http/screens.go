package http

import (
	"net/http"
)

func (rout *Router) GetData(w http.ResponseWriter, r *http.Request) {
	user := "default"
	code, resErr := rout.ser.ScreenSer.SGetImage(user)
	if code != 200 {
		http.Error(w, resErr, code)
	} else {
		w.Write([]byte(resErr))
	}
}
