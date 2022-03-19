package main

import (
	"fmt"

	"github.com/WeiWeiCheng123/Golang-Login-API/handler"
	"github.com/WeiWeiCheng123/Golang-Login-API/lib/config"
	"github.com/WeiWeiCheng123/Golang-Login-API/lib/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func init() {
	connectStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.GetStr("DB_HOST"), config.GetStr("DB_PORT"), config.GetStr("DB_NAME"),
		config.GetStr("DB_USERNAME"), config.GetStr("DB_PASSWORD"), config.GetStr("DB_SSL_MODE"))
	db, _ := xorm.NewEngine("postgres", connectStr)

	middleware.Init(db)
}

func engine() *gin.Engine {

	server := gin.Default()
	server.POST("/api/v1/signin", middleware.Plain(), handler.SingIn)
	server.POST("/api/v1/signup", middleware.Plain(), handler.SignUp)

	return server
}

func main() {
	server := engine()
	server.Run(":8080")
}
