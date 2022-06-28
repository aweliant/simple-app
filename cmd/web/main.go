package main

import (
	"fmt"
	"github.com/aweliant/simple-app/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	//http.ResponseWriter 是一种 io.Writer
	//This is a handler func
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello world!")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("Bytes written:" + string(n))
	//})
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
