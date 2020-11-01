package controller

import "net/http"

//Index view Index
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`111`))
}
