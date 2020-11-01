package controller

import (
	"io/ioutil"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	buf, _ := ioutil.ReadFile("./views/list.html")
	w.Write(buf)
}
