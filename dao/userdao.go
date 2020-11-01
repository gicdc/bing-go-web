package dao

import (
	"github.com/gicdc/bing-go-web/model"
	"github.com/gicdc/bing-go-web/utils"
)

//GetUsers 获取所有的用户信息
func GetUsers() ([]*model.User, error) {
	users := []*model.User{}
	query := "select * from users"
	err := utils.Db.Select(users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
