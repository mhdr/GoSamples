package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/index", serverFile)
	http.ListenAndServe(":8000", nil)
}

func sayHello(writer http.ResponseWriter, request *http.Request)  {
	io.WriteString(writer,"Hello World")
}

func serverFile(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter,request,"index.html")
}