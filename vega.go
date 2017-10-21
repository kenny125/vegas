/*

 __      __
 \ \    / /
  \ \  / /__  __ _  __ _
   \ \/ / _ \/ _` |/ _` |
    \  /  __/ (_| | (_| |
     \/ \___|\__, |\__,_|
              __/ |
             |___/

           ,ggg,                                                   ,ggg,         gg
          dP""8I           8I         I8                          dP""Y8a        88                                   ,dPYb,
         dP   88           8I         I8                          Yb, `88        88                                   IP'`Yb
        dP    88           8I   gg 88888888                        `"  88        88                                   I8  8I
       ,8'    88           8I   ""    I8                               88        88                                   I8  8'
       d88888888     ,gggg,8I   gg    I8   gg     gg    ,gggg,gg       88        88  gg,gggg,   gg,gggg,     ,gggg,gg I8 dP
 __   ,8"     88    dP"  "Y8I   88    I8   I8     8I   dP"  "Y8I       88        88  I8P"  "Yb  I8P"  "Yb   dP"  "Y8I I8dP
dP"  ,8P      Y8   i8'    ,8I   88   ,I8,  I8,   ,8I  i8'    ,8I       88        88  I8'    ,8i I8'    ,8i i8'    ,8I I8P
Yb,_,dP       `8b,,d8,   ,d8b,_,88,_,d88b,,d8b, ,d8I ,d8,   ,d8b,      Y8b,____,d88,,I8 _  ,d8',I8 _  ,d8',d8,   ,d8b,d8b,_
 "Y8P"         `Y8P"Y8888P"`Y88P""Y88P""Y8P""Y88P"888P"Y8888P"`Y8       "Y888888P"Y8PI8 YY88888PI8 YY88888P"Y8888P"`Y8P'"Y88
                                                ,d8I'                                I8         I8
                                              ,dP'8I                                 I8         I8
                                             ,8"  8I                                 I8         I8
                                             I8   8I                                 I8         I8
                                             `8, ,8I                                 I8         I8
                                              `Y8P"                                  I8         I8

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
