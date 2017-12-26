package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index is root path.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!")
}

// ItemIndex return item lists.
func ItemIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Item Index!")
}

// ItemShow return a item.
func ItemShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Item show: %s", ps.ByName("itemId"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/items", ItemIndex)
	router.GET("/items/:itemId", ItemShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
