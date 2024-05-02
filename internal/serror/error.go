package serror

import "errors"

var (
	WrongCaptchaError  = errors.New("wrong captcha")
	UserNameExistError = errors.New("userName already exist")
	WrongPassWordError = errors.New("password wrong")
	NoSuchUserError    = errors.New("no such user")
	WrongRangeError    = errors.New("wrong range error")
)
