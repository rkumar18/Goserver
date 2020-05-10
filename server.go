package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)



func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/about", aboutus)
	router.HandleFunc("/more", learnmore)
	router.HandleFunc("/contact", contactus)
	router.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8001",router))
}
func index(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("pages/index.html")
	fmt.Fprintf(w, string(data))
}
func aboutus(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("pages/about.html")
	fmt.Fprintf(w, string(data))
}
func learnmore(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("pages/more.html")
	fmt.Fprintf(w, string(data))
}
func contactus(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("pages/contact.html")
	fmt.Fprintf(w, string(data))
}