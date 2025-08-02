package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myblok/apisrc"
	"myblok/mylog"
	"myblok/sqlmodel"
)

func InitDB() *gorm.DB {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/myblok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		mylog.Mlogger.Println(err)
		return nil
	}
	return db
}

func main() {

	mylog.Initlog()
	db := InitDB()
	r := gin.Default()
	// generate users posts comments table
	sqlmodel.Run(db)

	err := apisrc.RegisterAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}

	token, err2 := apisrc.LoginAPI(r, db)
	if err2 != nil {
		mylog.Mlogger.Println(err)
	}

	err = apisrc.CreateArticleAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}
	err = apisrc.UpdateArticleAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}
	err = apisrc.DeleteArticleAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}
	err = apisrc.ReadORListArticleAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}
	err = apisrc.CreateCommentAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}
	err = apisrc.ListCommentAPI(r, db)
	if err != nil {
		mylog.Mlogger.Println(err)
	}
	mylog.Mlogger.Println(token)
	r.Run(":8080")
	return
}
