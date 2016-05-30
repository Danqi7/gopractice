package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"io"
	"strconv"
	"crypto/md5"
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
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		for k, v := range r.Form {
			fmt.Println("key: ", k)
			fmt.Println("val: ", strings.Join(v, " AND "))
		}
	
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, token)

	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		fmt.Println("token: ", token)
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
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


