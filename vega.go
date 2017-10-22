package main

import "fmt"
import "time"
import "net/http"
import (
	"io"
	"log"
)

/*

 */
const PORT = ":8000"

var mux map[string]func(http.ResponseWriter, *http.Request)

/*

                       _ _
  /\  /\__ _ _ __   __| | | ___ _ __ ___
 / /_/ / _` | '_ \ / _` | |/ _ \ '__/ __|
/ __  / (_| | | | | (_| | |  __/ |  \__ \
\/ /_/ \__,_|_| |_|\__,_|_|\___|_|  |___/


*/

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "The page at"+r.URL.String()+" is not implemented yet.")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is neat")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, tm)
}
func serverInitalize() { // includes all handleFunc functions
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = indexHandler
	mux["/about"] = aboutHandler
	mux["/hello"] = helloHandler
	mux["/time"] = timeHandler
}

/*

              _
  /\/\   __ _(_)_ __
 /    \ / _` | | '_ \
/ /\/\ \ (_| | | | | |
\/    \/\__,_|_|_| |_|


*/

func main() {
	server := http.Server{
		Addr:    PORT,
		Handler: &myHandler{},
	}

	fmt.Println("Server started at ", time.Now())
	go serverInitalize() // Initialize all handlers as a goroutine
	log.Fatal(server.ListenAndServe())
}
