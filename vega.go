/*


 ██▒   █▓▓█████   ▄████  ▄▄▄        ██████ 
▓██░   █▒▓█   ▀  ██▒ ▀█▒▒████▄    ▒██    ▒ 
 ▓██  █▒░▒███   ▒██░▄▄▄░▒██  ▀█▄  ░ ▓██▄   
  ▒██ █░░▒▓█  ▄ ░▓█  ██▓░██▄▄▄▄██   ▒   ██▒
   ▒▀█░  ░▒████▒░▒▓███▀▒ ▓█   ▓██▒▒██████▒▒
   ░ ▐░  ░░ ▒░ ░ ░▒   ▒  ▒▒   ▓▒█░▒ ▒▓▒ ▒ ░
   ░ ░░   ░ ░  ░  ░   ░   ▒   ▒▒ ░░ ░▒  ░ ░
     ░░     ░   ░ ░   ░   ░   ▒   ░  ░  ░  
      ░     ░  ░      ░       ░  ░      ░  
     ░                                     


*/

package main

import "fmt"
import "time"
import "net/http"
import "log"

/*

                       _ _
  /\  /\__ _ _ __   __| | | ___ _ __ ___
 / /_/ / _` | '_ \ / _` | |/ _ \ '__/ __|
/ __  / (_| | | | | (_| | |  __/ |  \__ \
\/ /_/ \__,_|_| |_|\__,_|_|\___|_|  |___/


*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is neat")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func serverInitalize() { // includes all handleFunc functions
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
}

/*

              _
  /\/\   __ _(_)_ __
 /    \ / _` | | '_ \
/ /\/\ \ (_| | | | | |
\/    \/\__,_|_|_| |_|


*/

func main() {
	const PORT = ":8000"
	fmt.Println("Server started at ", time.Now())
	go serverInitalize() // Initialize all handlers as a goroutine
	log.Fatal(http.ListenAndServe(PORT, nil))
}
