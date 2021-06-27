package http

import (
	"fmt"
	"net/http"
)

func (rout *Router) GetData(w http.ResponseWriter, r *http.Request) {
	user := "default"
	fmt.Println("OK")
	code, resErr := rout.ser.ScreenSer.SGetImage(user)
	if code != 200 {
		http.Error(w, resErr, code)
	} else {
		w.Write([]byte(resErr))
	}
}
