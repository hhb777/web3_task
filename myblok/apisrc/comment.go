package apisrc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myblok/mylog"
	"myblok/sqlmodel"
	"net/http"
)

func CreateCommentAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.PUT("/comment", func(c *gin.Context) {
		type tcomment struct {
			sqlmodel.Comment
			Token string `json:"token" ;binding:"required"`
		}
		mcomment := tcomment{}
		err = c.ShouldBindJSON(&mcomment)
		if err != nil {
			mylog.Mlogger.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if ok, err1 := TokenAuth(mcomment.Token); !ok {
			mylog.Mlogger.Println(err1.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": err1.Error()})
			err = err1
			return
		}

		res := db.Create(&mcomment.Comment)
		if res.Error != nil {
			mylog.Mlogger.Println(res.Error.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "comment created successfully"})
		mylog.Mlogger.Println("comment created successfully")
	})
	return
}
func ListCommentAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.GET("/comment", func(c *gin.Context) {
		type tcomment struct {
			Post_id uint `json:"post_id" ;query:"post_id"`
			User_id uint `json:"user_id" ;query:"user_id"`
		}
		qcomment := tcomment{}
		err = c.ShouldBindQuery(&qcomment)
		if err != nil {
			mylog.Mlogger.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var comments []sqlmodel.Comment
		res := db.Where(sqlmodel.Comment{Post_id: qcomment.Post_id, User_id: qcomment.User_id}).Find(&comments)
		if res.Error != nil {
			mylog.Mlogger.Println(res.Error.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"children": comments})
		mylog.Mlogger.Println("comments list successfully")
	})
	return
}
