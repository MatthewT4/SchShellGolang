package http

import (
	"encoding/json"
	"fmt"
	"github.com/MatthewT4/SchShellGolang/internal/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Router struct {
	ser *service.Service
}

func NewRouter(db *mongo.Database) *Router {
	return &Router{service.NewService(db)}
}
func (rout *Router) AddCatalog(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	cat := &service.Catalog{}
	d.Decode(cat)
	fmt.Fprint(w, cat.Name+" "+string(cat.Type))
}
func (rout *Router) Start() {
	r := mux.NewRouter()
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/login", LoginUser)
	r.HandleFunc("/addcatalog", rout.AddCatalog).Methods("POST") //POST
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
