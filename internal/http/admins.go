package http

import (
	"encoding/json"
	"fmt"
	"github.com/MatthewT4/SchShellGolang/internal/service"
	"io/ioutil"
	"log"
	"net/http"
)

func (rout *Router) AddCatalog(w http.ResponseWriter, r *http.Request) {
	cat := service.Catalog{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &cat)
	if err != nil {
		panic(err)
	}
	log.Println(cat.Name, cat.Type)
	userGroup := r.Context().Value(GroupKey)
	if userGroup.(int) < service.Administration || cat.Holder == "" { // you can add catalogues for other users only from the "administrator" level
		userInter := r.Context().Value(UserKey)
		cat.Holder = userInter.(string)
	}
	code, _ := rout.ser.CatalogSer.SAddCatalog(cat)
	if code == 200 {
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Error", code)
	}
}
func (rout *Router) GetCatalogs(w http.ResponseWriter, r *http.Request) {
	userInter := r.Context().Value(UserKey)
	holder := userInter.(string)
	data, err := rout.ser.CatalogSer.SGetCatalogs(holder)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		var OnlyGetCatalogsHTTP struct {
			NameCatalogs []string `bson:"name_catalogs"`
		}
		OnlyGetCatalogsHTTP.NameCatalogs = data
		res, _ := json.Marshal(OnlyGetCatalogsHTTP)
		w.Write(res)
	}

}

func (rout *Router) GetDataInCatalog(w http.ResponseWriter, r *http.Request) {
	holder := r.URL.Query().Get("holder")
	userGroup := r.Context().Value(GroupKey)
	if userGroup.(int) < service.Administration || holder == "" { // you can add catalogues for other users only from the "administrator" level
		userInter := r.Context().Value(UserKey)
		holder = userInter.(string)
	}
	catName := r.URL.Query().Get("name")
	log.Println(catName, holder)
	code, res := rout.ser.CatalogSer.SGetDataInCatalog(holder, catName)
	if code != 200 {
		http.Error(w, "error", code)
	} else {
		w.Write(res)
	}
}

func (rout *Router) InsertDataInCatalog(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var OnlyInsertDataVar struct {
		NameCatalog string `json:"name_catalog"`
		Data        string `json:"data"`
	}
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &OnlyInsertDataVar)
	userInter := r.Context().Value(UserKey)
	user := fmt.Sprint(userInter)
	code, mes := rout.ser.CatalogSer.SInsertDataInCatalog(user, OnlyInsertDataVar.NameCatalog, OnlyInsertDataVar.Data)
	if code != 200 {
		http.Error(w, mes, code)
	} else {
		w.Write([]byte(mes))
	}
}
