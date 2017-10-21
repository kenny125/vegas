/*

 __      __
 \ \    / /
  \ \  / /__  __ _  __ _
   \ \/ / _ \/ _` |/ _` |
    \  /  __/ (_| | (_| |
     \/ \___|\__, |\__,_|
              __/ |
             |___/

*/

package main

import "fmt"
import "time"
import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is neat")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func main() {
	fmt.Println("Server started at ", time.Now())

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)

}
