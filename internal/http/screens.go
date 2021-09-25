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

func (rout *Router) NewScreen(w http.ResponseWriter, r *http.Request) {
	ret_code, time, screen_token, screen_code, err := rout.ser.ScreenSer.NewScreen()
	if ret_code != 200 || err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	cookie := http.Cookie{Name: "authToken", Value: screen_token, Expires: time}
	http.SetCookie(w, &cookie)
	w.Write([]byte(screen_code))
}
