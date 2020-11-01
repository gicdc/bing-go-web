package main

import (
	"fmt"
	"net/http"

	"github.com/gicdc/bing-go-web/controller"
)

func main() {

	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/list", controller.List)
	http.HandleFunc("/getUsers", controller.GetUsers)
	http.ListenAndServe(":8080", nil)
	fmt.Println("hello")
}
