package mysql

import (
	"cmdTest/internal/database"
	"cmdTest/internal/dto/model"
	"time"
)

var db = database.DB

type UserDAOMysql struct{}

func (m *UserDAOMysql) CreateUser(user *model.User) error {
	sqlStr := "INSERT INTO users(uid, name, role, user_name, pass_word, email) VALUE (:uid, :name, :role, :user_name, :pass_word, :email)"
	_, err := db.NamedExec(sqlStr, user)
	return err
}

func (m *UserDAOMysql) GetCaptcha(userName string) (captcha string, err error) {
	sqlStr := "SELECT token FROM captcha WHERE username = ? AND begin > ? ORDER BY begin DESC LIMIT 1"
	err = db.Get(&captcha, sqlStr, userName, time.Now().Unix()-5*60)
	return
}

func (m *UserDAOMysql) SaveCaptcha(userName, captcha string) error {
	sqlStr := "INSERT INTO captcha(username, token, begin) VALUE (?, ?, ?)"
	_, err := db.Exec(sqlStr, userName, captcha, time.Now().Unix())
	return err
}

func (m *UserDAOMysql) GetUserNameCnt(userName string) (cnt int, err error) {
	sqlStr := "SELECT COUNT(1) FROM users WHERE user_name = ?"
	err = db.Get(&cnt, sqlStr, userName)
	return
}

func (m *UserDAOMysql) GetUserByUserName(userName string) (user *model.User, err error) {
	user = &model.User{}
	sqlStr := "SELECT * FROM users WHERE user_name = ?"
	err = db.Get(user, sqlStr, userName)
	return
}

func (m *UserDAOMysql) GetUserByUID(uid uint) (user *model.User, err error) {
	user = &model.User{}
	sqlStr := "SELECT * FROM users WHERE uid = ?"
	err = db.Get(user, sqlStr, uid)
	return
}

func (m *UserDAOMysql) UpdateUser(name, newPassWord string, uid uint) error {
	sqlStr := "UPDATE users SET name = ?, pass_word = ? WHERE uid = ?"
	_, err := db.Exec(sqlStr, name, newPassWord, uid)
	return err
}

func (m *UserDAOMysql) GetAllUser() (users []model.User, err error) {
	sqlStr := "SELECT * FROM users"
	err = db.Select(&users, sqlStr)
	return
}

func (m *UserDAOMysql) DeleteUser(uid string) error {
	sqlStr := "DELETE FROM users WHERE uid = ?"
	_, err := db.Exec(sqlStr, uid)
	return err
}

func (m *UserDAOMysql) ReSetUser(uid, passWord string) error {
	sqlStr := "UPDATE users SET pass_word = ? WHERE uid = ?"
	_, err := db.Exec(sqlStr, passWord, uid)
	return err
}
