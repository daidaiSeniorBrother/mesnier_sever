package service

import (
	"mesnier/confs"
	"mesnier/model"
	"mesnier/utils"
	"strings"
)

type UserService struct {
}

func (u *UserService) UserLogin(user *model.User) (token string, err error) {
	err = confs.DB.Where("user_account", user.UserAccount).Where("user_password", user.UserPassword).First(&user).Error
	if user.UserId <= 0 {
		err = utils.NewBusinessError(-1, "登录用户账号或密码错误")
	} else {
		token = strings.ToUpper(utils.UUID())
	}
	return
}

func (u *UserService) UserList(*model.User) (pageResponse model.PageResponse, err error) {
	var (
		users []model.User
		count int64
	)
	err = confs.DB.Where("delete_flag", 0).Find(&users).Count(&count).Error
	pageResponse.ObjectData = users
	pageResponse.Total = count
	if len(users) <= 0 {
		err = utils.NewBusinessError(-1, "未找到用户数据")
	}
	return
}
