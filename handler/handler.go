package handler

import (
	"errors"

	"net/http"

	"github.com/WeiWeiCheng123/Golang-Login_system/lib/constant"
	"github.com/WeiWeiCheng123/Golang-Login_system/lib/function"
	"github.com/WeiWeiCheng123/Golang-Login_system/model"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"golang.org/x/crypto/bcrypt"
)

var user struct {
	username string
	password string
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
	result, err := db.Where("username= ?", user.username).Get(&user_tab)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	if result == false {
		c.Set(constant.StatusCode, http.StatusUnauthorized)
		c.Set(constant.Error, "user not exist or wrong password")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user_tab.Passwd), []byte(user.password)); err != nil {
		c.Set(constant.StatusCode, http.StatusUnauthorized)
		c.Set(constant.Error, "user not exist or wrong password")
		return
	}

	out := "Welcome, " + user.username
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

	if err := function.CheckUserIsAccept(user.username); err != nil {
		c.Set(constant.StatusCode, http.StatusForbidden)
		c.Set(constant.Error, err)
		return
	}

	hashed_password, err := function.HashPassword(user.password)
	if err != nil {
		c.Set(constant.StatusCode, http.StatusInternalServerError)
		c.Set(constant.Error, err)
		return
	}

	q := `INSERT INTO user(username, password) VALUE($1,$2) WHERE NOT EXIST(SELECT 1 FROM user WHERE username = $1)`
	db := c.MustGet(constant.DB).(*xorm.Engine)
	res, err := db.Exec(q, user.username, hashed_password)
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

	out := "User " + user.username + "created"
	c.Set(constant.StatusCode, http.StatusCreated)
	c.Set(constant.Output, out)
}
