package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gicdc/bing-go-web/dao"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := dao.GetUsers()
	buf, _ := json.Marshal(users)
	w.Write(buf)
}
