package middleware

import (
	"fmt"

	"github.com/WeiWeiCheng123/Golang-Login-API/lib/constant"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func Init(database *xorm.Engine) {
	db = database
}

func Plain() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("header= ",c.GetHeader("Authorization"))
		c.Set(constant.DB, db)
		c.Set(constant.StatusCode, nil)
		c.Set(constant.Error, nil)
		c.Set(constant.Output, nil)
		c.Next()

		statusCode := c.GetInt(constant.StatusCode)
		err := c.MustGet(constant.Error)
		output := c.MustGet(constant.Output)
		if err != nil {
			c.String(statusCode, err.(error).Error())
		} else {
			c.String(statusCode, output.(string))
		}
	}
}

