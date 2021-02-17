package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/",sayHello)
	err:= http.ListenAndServe(":8080" , nil)
	if err != nil {
		log.Fatal("ListenAndServe: " , err)
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm() // 解析参数
	fmt.Println(request.Form) // 输出打印信息
	fmt.Printf("Path: %s \n Host: %s \n" , request.URL.Path , request.Host)
	for k, v := range request.Form  {
		fmt.Printf("key : %s \n value: %s \n " , k,strings.Join(v,""))
	}
	_,_ = fmt.Fprintf(writer,"Hello Web , %s !" , request.Form.Get("name"))

}
