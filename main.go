package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/login" , login)
	err:= http.ListenAndServe(":8080" , nil)
	if err != nil {
		log.Fatal("ListenAndServe: " , err)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method: ",request.Method)
	m := request.Method
	if m == "GET" {
		// get
		t,_ := template.ParseFiles("login.tpl")
		log.Println(t.Execute(writer,nil))
	} else {
		// login
		_ = request.ParseForm() // 解析参数
		fmt.Println("username: ",request.Form["username"])
		fmt.Println("password: ",request.Form["password"])
		if pwd := request.Form.Get("password") ; pwd == "123456" {
			fmt.Fprintf(writer,"welcome login ,Hello %s" , request.Form.Get("username"))
		} else {
			fmt.Fprintf(writer,"error login , maybe password wrong , please login again.")
		}
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm() // 解析参数
	fmt.Println(request.Form) // 输出打印信息
	fmt.Printf("Path: %s \nHost: %s \n" , request.URL.Path , request.Host)
	for k, v := range request.Form  {
		fmt.Printf("key : %s \nvalue: %s \n " , k,strings.Join(v,""))
	}
	_,_ = fmt.Fprintf(writer,"Hello Web , %s !" , request.Form.Get("name"))

}
