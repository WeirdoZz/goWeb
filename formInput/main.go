package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数
	fmt.Println(r.Form) //输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}

	fmt.Fprintf(w, "hello weirdo")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./login.tmpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm:", err)
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("username:", r.Form["password"])
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
