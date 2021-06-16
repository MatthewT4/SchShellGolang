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
	r := mux.NewRouter()
	//r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/login", LoginUser)
	r.HandleFunc("/addcatalog", rout.AddCatalog).Methods("POST") //POST
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/admin", r)

	rou := mux.NewRouter()
	rou.HandleFunc("/getdata", rout.GetData)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
