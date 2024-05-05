package api

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/response"
	"cmdTest/internal/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Register
//
//	@Summary	用户注册
//	@Tags		User
//	@Param		ud	body	model.UserJson	true	"用户信息"
//	@Router		/api/v1/user/new [post]
func Register(c *gin.Context) {
	var userJson model.UserJson
	err := c.ShouldBindJSON(&userJson)
	if err != nil {
		response.Fail(c.Writer, fmt.Sprintf("data bind serror, err:%s", err.Error()), 400)
		return
	}

	err = server.NewUser(userJson)

	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}
	response.Success(c.Writer, "success", nil)
}

// Verify
//
//	@Summary	发送验证码
//	@Tags		User
//	@Param		userName	formData	string	true	"用户名"
//	@Param		email		formData	string	true	"邮箱"
//	@Router		/api/v1/user/v [post]
func Verify(c *gin.Context) {
	userName := c.PostForm("userName")
	email := c.PostForm("email")

	err := server.SendCaptcha(userName, email)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Fail(c.Writer, "success", 200)
}

// CheckUserName
//
//	@Summary	检验用户名是否重复
//	@Tags		User
//	@Param		userName	query	string	true	"用户名"
//	@Router		/api/v1/user/check [get]
func CheckUserName(c *gin.Context) {
	userName := c.Query("userName")

	err := server.CheckUserName(userName)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// Login
//
//	@Summary	登录
//	@Tags		User
//	@Param		userName	formData	string	true	"用户名"
//	@Param		passWord	formData	string	true	"密码"
//	@Router		/api/v1/user/login [post]
func Login(c *gin.Context) {
	userName := c.PostForm("userName")
	passWord := c.PostForm("passWord")

	token, err := server.UserLogin(userName, passWord)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", token)
}

// UserRevise
//
//	@Summary	用户修改信息
//	@Tags		User
//	@Param		name			formData	string	false	"用户名称"
//	@Param		oldPassWord		formData	string	false	"旧密码"
//	@Param		newPassWord		formData	string	false	"新密码"
//	@Param		Authorization	header		string	true	"token"
//
//	@Router		/api/v1/user/ru [post]
func UserRevise(c *gin.Context) {
	name := c.PostForm("name")
	oldPassWord := c.PostForm("oldPassWord")
	newPassWord := c.PostForm("newPassWord")
	token := c.GetHeader("Authorization")

	err := server.UpdateUser(name, oldPassWord, newPassWord, token)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// UserInfo
//
//	@Summary	获取当前登录用户信息
//	@Tags		User
//	@Param		Authorization	header	string	true	"token"
//	@Router		/api/v1/user/info [get]
func UserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")

	user, err := server.UserInfo(token)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", *user)
}

// GetAllUser
//
//	@Summary	获取全部用户信息
//	@Tags		User
//	@Router		/api/v1/user/all [get]
func GetAllUser(c *gin.Context) {
	users, err := server.GetAllUser()
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", users)
}

// DeleteUser
//
//	@Summary	移除用户
//	@Tags		User
//	@Param		uid	formData	string	true	"uid"
//	@Router		/api/v1/user/d [post]
func DeleteUser(c *gin.Context) {
	uid := c.PostForm("uid")

	err := server.DeleteUser(uid)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}

// ReSetUser
//
//	@Summary	重置用户密码
//	@Tags		User
//	@Param		uid	formData	string	true	"uid"
//	@Router		/api/v1/user/rp [post]
func ReSetUser(c *gin.Context) {
	uid := c.PostForm("uid")

	err := server.ReSetUser(uid)
	if err != nil {
		response.Fail(c.Writer, err.Error(), 500)
		return
	}

	response.Success(c.Writer, "success", nil)
}
