package dao

import "cmdTest/internal/dto/model"

type UserDao interface {
	CreateUser(user *model.User) error
	GetCaptcha(userName string) (string, error)
	SaveCaptcha(userName, captcha string) error
	GetUserNameCnt(userName string) (int, error)
	GetUserByUserName(userName string) (*model.User, error)
	GetUserByUID(uid uint) (*model.User, error)
	UpdateUser(name, newPassWord string, uid uint) error
	GetAllUser() ([]model.User, error)
	DeleteUser(uid string) error
	ReSetUser(uid, passWord string) error
}
