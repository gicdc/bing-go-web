package main

import (
	"fmt"
	"net/http"

	"github.com/gicdc/bing-go-web/controller"
)

func main() {

	http.HandleFunc("/index", controller.Index)
	http.ListenAndServe(":8080", nil)
	fmt.Println("hello")
}
