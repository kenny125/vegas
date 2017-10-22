package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/*
PORT on which to server. Default = 666
*/
const PORT = ":666"

var mux map[string]func(http.ResponseWriter, *http.Request)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf(" %s took %d nanoseconds", name, elapsed.Nanoseconds())
}

/*

                       _ _
  /\  /\__ _ _ __   __| | | ___ _ __ ___
 / /_/ / _` | '_ \ / _` | |/ _ \ '__/ __|
/ __  / (_| | | | | (_| | |  __/ |  \__ \
\/ /_/ \__,_|_| |_|\__,_|_|\___|_|  |___/


*/

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "myHandler")
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "The page at"+r.URL.String()+" is not implemented yet.")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "hello")
	io.WriteString(w, "Hello world!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "indexHandler")
	io.WriteString(w, "VEGAS is LIVE")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "aboutHandler")
	fmt.Fprintf(w, "This is VEGAS")

}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "timeHandler")
	fmt.Fprintf(w, time.Now().Format(time.RFC1123))
}

func serverInitalize() { // includes all handleFunc functions
	defer timeTrack(time.Now(), "serverInitalize")
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

	fmt.Println("Server started at ", time.Now().Format(time.RFC1123))
	go serverInitalize() // Initialize all handlers as a goroutine
	log.Fatal(server.ListenAndServe())
}
