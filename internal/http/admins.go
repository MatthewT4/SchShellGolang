package http

import (
	"encoding/json"
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
	code, addErr := rout.ser.CatalogSer.SAddCatalog(cat)
	if code != 200 {
		http.Error(w, addErr.Error(), code)
	} else {
		w.Write([]byte("OK"))
	}
}
