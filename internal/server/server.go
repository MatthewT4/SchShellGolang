package server

import (
	"fmt"
	"net/http"
	"time"
)
type Config struct {
	Addr string
}
func Test(n string) {
	fmt.Println(n)
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	time.Sleep(10*time.Second)
	fmt.Fprintf(w, "444444444444333333333333333333OKOKOKKOKOKOKOKOKOKOKO")
}
func Start(conf Config) {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/getData", mainHandler)
	http.ListenAndServe(conf.Addr, nil)
}