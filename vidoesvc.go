package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func handleVideo(w http.ResponseWriter, req *http.Request) {

}

func videosvc() {
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleVideo)

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":3000", n)
}
