package handler

import (
	"errors"

	"net/http"

	"github.com/WeiWeiCheng123/Golang-Login-API/lib/constant"
	"github.com/WeiWeiCheng123/Golang-Login-API/lib/function"
	"github.com/WeiWeiCheng123/Golang-Login-API/model"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var user struct {
	Username string `json:"username"`
	Password string `json:"passwd"`
}

func SingIn(c *gin.Context) {

	err := c.BindJSON(&user)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusBadRequest)
		c.Set(constant.Error, err)
		return
	}

	user_tab := model.User{}
	db := c.MustGet(constant.DB).(*xorm.Engine)
	result, err := db.Where("username= ?", user.Username).Get(&user_tab)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	if result == false {
		c.Set(constant.StatusCode, http.StatusUnauthorized)
		c.Set(constant.Error, errors.New("user not exist or wrong password"))
		return
	}
	if err = function.Compare(user_tab.Passwd, user.Password); err != nil {
		c.Set(constant.StatusCode, http.StatusUnauthorized)
		c.Set(constant.Error, errors.New("user not exist or wrong password"))
		return
	}

	out := "Welcome, " + user.Username
	c.Set(constant.StatusCode, http.StatusOK)
	c.Set(constant.Output, out)
}

func SignUp(c *gin.Context) {
	err := c.BindJSON(&user)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusBadRequest)
		c.Set(constant.Error, err)
		return
	}

	if err := function.CheckUserIsAccept(user.Username); err != nil {
		c.Set(constant.StatusCode, http.StatusForbidden)
		c.Set(constant.Error, err)
		return
	}

	hashed_password, err := function.HashPassword(user.Password)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	q := `INSERT INTO user_table(username, passwd) SELECT ?,? WHERE NOT EXISTS (SELECT 1 FROM user_table WHERE username = ?)`
	db := c.MustGet(constant.DB).(*xorm.Engine)
	res, err := db.Exec(q, user.Username, hashed_password, user.Username)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	affected, err := res.RowsAffected()
	if affected == 0 {
		c.Set(constant.StatusCode, http.StatusForbidden)
		c.Set(constant.Error, errors.New("the username is already used"))
		return
	}

	out := "User " + user.Username + " created"
	c.Set(constant.StatusCode, http.StatusCreated)
	c.Set(constant.Output, out)
}
