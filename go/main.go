package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8080", "Port to host the application on")
}

func main() {
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		fmt.Fprintln(w, "Hello", name)
	}).Methods(http.MethodGet)

	fmt.Println("Serving on port:", port)

	log.Fatalln(http.ListenAndServe(net.JoinHostPort("", port), router))
}
