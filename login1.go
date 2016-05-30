package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func helloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " AND "))
	}
	fmt.Fprintf(w, "hello" + r.URL.Path[1:])
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		for k, v := range r.Form {
			fmt.Println("key: ", k)
			fmt.Println("val: ", strings.Join(v, " AND "))
		}
	
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)

	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}

}

func main() {
	http.HandleFunc("/", helloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}


