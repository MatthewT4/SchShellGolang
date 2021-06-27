package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserK string

const (
	UserKey  UserK  = "Login"
	GroupKey string = "Group"
)

func (rout *Router) Authorization(w http.ResponseWriter, r *http.Request) {
	var DataLogin struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &DataLogin)
	if err != nil {
		http.Error(w, "Error parse data", 403)
	}
	res, time, er := rout.ser.UserSer.Authorization(DataLogin.Login, DataLogin.Password)
	if er != nil {
		if res == "404" {
			http.Error(w, er.Error(), 404) // данные входа не корректны
			return
		}
		http.Error(w, er.Error(), 500)
	}
	//.Header().Add("authToken", res)
	cookie := http.Cookie{Name: "authToken", Value: res, Expires: time}
	http.SetCookie(w, &cookie)
	w.Write([]byte("OK"))
}

func (rout *Router) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, er := r.Cookie("authToken")
		if er != nil {
			http.Error(w, "not cookie", 401)
			return
		}
		group, pToken, err := rout.ser.UserSer.Authentication(cookie.Value)
		if err != nil {
			http.Error(w, "Authentication error"+err.Error(), 404)
		} else {
			ctx := context.WithValue(r.Context(), UserKey, pToken)
			r = r.WithContext(ctx)
			ctxt := context.WithValue(r.Context(), GroupKey, group)
			r = r.WithContext(ctxt)
			next.ServeHTTP(w, r)
		}
	})
}
