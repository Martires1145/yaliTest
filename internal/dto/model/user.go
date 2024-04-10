package model

type User struct {
	Uid      uint   `db:"uid" json:"uid,string"`
	Role     int    `db:"role" json:"role"`
	Name     string `db:"name" json:"name"`
	UserName string `db:"user_name" json:"userName"`
	PassWord string `db:"pass_word" json:"passWord"`
	Email    string `db:"email" json:"email"`
}

type UserJson struct {
	Role     int    `db:"role" json:"role"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Email    string `db:"email" json:"email"`
	Captcha  string `db:"captcha" json:"captcha"`
}

type Captcha struct {
	Id       uint   `db:"id"`
	Begin    uint   `db:"begin"`
	UserName string `db:"username"`
	Token    string `db:"token"`
}

func (j UserJson) ToUser() User {
	return User{
		Name:     j.Name,
		UserName: j.UserName,
		PassWord: j.PassWord,
		Email:    j.Email,
		Role:     j.Role,
	}
}
