package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func helloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
 	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Danqi!")	

}


func main() {
	http.HandleFunc("/", helloName)
	err := http.ListenAndServe(":9090", nil)
	if  err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

}
