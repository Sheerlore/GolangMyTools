//Goで作る基本的なhttpサーバー
//> go run server.go http://localhost:8081/
//で起動する 

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

//func echoString(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello")
//}

func incrementCounter(w http.ResponseWriter, r *http.Request){
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServeTLS(":443", "./server.crt", "./server.key", nil))
}
