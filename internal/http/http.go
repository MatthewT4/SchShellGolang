package http

import (
	"github.com/MatthewT4/SchShellGolang/internal/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Router struct {
	ser *service.Service
}

func NewRouter(db *mongo.Database) *Router {
	return &Router{service.NewService(db)}
}

func (rout *Router) Start() {
	rou := mux.NewRouter()
	r := rou.PathPrefix("/api/").Subrouter()
	r.HandleFunc("/getdata", rout.GetData)
	r.HandleFunc("/login", rout.Authorization)
	r.HandleFunc("/screen_register", rout.NewScreen)
	rou.Handle("/", r)

	rAdm := rou.PathPrefix("/api/admin").Subrouter()
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/login", LoginUser)
	rAdm.HandleFunc("/addcatalog", rout.AddCatalog).Methods("POST")
	rAdm.HandleFunc("/insertdata", rout.InsertDataInCatalog).Methods("POST")
	rAdm.HandleFunc("/getcatalogs", rout.GetCatalogs).Methods("GET")
	rAdm.HandleFunc("/getdata", rout.GetDataInCatalog).Methods("GET")
	//r.HandleFunc("/articles", ArticlesHandler)
	rAdm.Use(rout.Authentication)

	srv := &http.Server{
		Handler: rou,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
