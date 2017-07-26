package main

import (
	"io"
	"net/http"
)

func main() {

	mux :=http.NewServeMux()
	mux.HandleFunc("/",sayHello)
	http.ListenAndServe(":8000", mux)
}

func sayHello(writer http.ResponseWriter, request *http.Request)  {
	io.WriteString(writer,"Hello World")
}