package service

import (
	"github.com/jikebang/Go-000/Week02/dao"
	"github.com/jikebang/Go-000/Week02/util"
)

type userService struct {
}


var UserService = new(userService)

func (u *userService) One(id string) (dao.User, error) {
	user := dao.User{
		Id: id,
	}
	return dao.One(user)
}

func (u *userService) Rows() (*util.PageRes, error) {
	users := make([]*dao.User, 0)
	return dao.Rows(users)
}

