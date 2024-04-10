package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"cmdTest/internal/serror"
	"cmdTest/pkg/util"
	"github.com/spf13/viper"
)

var userDao = &mysql.UserDAOMysql{}

func NewUser(userJson model.UserJson) error {
	err := checkCaptcha(userJson.UserName, userJson.Captcha)
	if err != nil {
		return err
	}

	uid := util.GetUID()

	user := userJson.ToUser()
	user.Uid = uid

	err = userDao.CreateUser(&user)
	return err
}

func checkCaptcha(userName string, captcha string) error {
	captGet, err := userDao.GetCaptcha(userName)
	if err != nil {
		return err
	}

	if captGet != captcha {
		return serror.WrongCaptchaError
	}
	return nil
}

func SendCaptcha(userName, email string) error {
	captcha, err := util.SendMessage(email)
	if err != nil {
		return err
	}

	err = userDao.SaveCaptcha(userName, captcha)
	return err
}

func CheckUserName(userName string) error {
	cnt, err := userDao.GetUserNameCnt(userName)
	if err != nil {
		return err
	}
	if cnt != 0 {
		return serror.UserNameExistError
	}

	return nil
}

func UserLogin(userName, passWord string) (token string, err error) {
	user, err := userDao.GetUserByUserName(userName)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", serror.NoSuchUserError
	}

	if user.PassWord != passWord {
		return "", serror.WrongPassWordError
	}

	token, err = util.CreatToken(user.Uid, user.Role)
	return
}

func UpdateUser(name, old, nw, token string) error {
	_, claims, err := util.ParseToken(token)
	if err != nil {
		return err
	}
	uid := claims.UID

	user, err := userDao.GetUserByUID(uid)

	if old != "" && user.PassWord != old {
		return serror.WrongPassWordError
	} else if old == "" {
		nw = user.PassWord
	}

	if name == "" {
		name = user.Name
	}

	err = userDao.UpdateUser(name, nw, uid)
	return err
}

func UserInfo(token string) (user *model.User, err error) {
	_, claims, err := util.ParseToken(token)
	if err != nil {
		return nil, err
	}
	uid := claims.UID

	user, err = userDao.GetUserByUID(uid)
	return
}

func GetAllUser() ([]model.User, error) {
	users, err := userDao.GetAllUser()
	return users, err
}

func DeleteUser(uid string) error {
	err := userDao.DeleteUser(uid)
	return err
}

func ReSetUser(uid string) error {
	defaultPassWord := viper.GetString("user.default")
	err := userDao.ReSetUser(uid, defaultPassWord)
	return err
}
