package apisrc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"myblok/mylog"
	"myblok/sqlmodel"
	"net/http"
	"strconv"
)

func CreateArticleAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.PUT("/article/create", func(c *gin.Context) {
		type sbody struct {
			sqlmodel.Post
			Token string `json:"token" ;binding:"required"`
		}
		gbody := sbody{}
		err = c.ShouldBindJSON(&gbody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			mylog.Mlogger.Println(err)
			return
		}
		mylog.Mlogger.Println(gbody)
		if ok, err1 := TokenAuth(gbody.Token); !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			mylog.Mlogger.Println(err1, "token:", gbody.Token)
			err = err1
			return
		}
		err = db.Create(&gbody.Post).Error
		if err != nil {
			mylog.Mlogger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Create article successfully",
		})
		mylog.Mlogger.Println("Create article successfully")
	})
	return
}

func ReadORListArticleAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.GET("/article/list", func(c *gin.Context) {
		lpost := []sqlmodel.Post{}
		res := db.Find(&lpost)
		if res.Error != nil {
			mylog.Mlogger.Println(res.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			err = res.Error
			return
		}
		c.JSON(http.StatusOK, gin.H{"children": lpost})
		mylog.Mlogger.Println("list article successfully")

	})
	r.GET("/article", func(c *gin.Context) {
		type QueryParam struct {
			Title  string `json:"title,omitempty" ;query:"title"`
			Userid string `json:"userid,omitempty" ;query:"userid"`
		}
		tquery := QueryParam{}
		err = c.ShouldBindQuery(&tquery)
		if err != nil {
			mylog.Mlogger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		int_userid, _ := strconv.ParseUint(tquery.Userid, 10, 0)
		mpost := sqlmodel.Post{}
		res := db.Where(&sqlmodel.Post{Title: tquery.Title, User_id: uint(int_userid)}).Preload(clause.Associations).Find(&mpost)
		if res.Error != nil {
			mylog.Mlogger.Println(res.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			err = res.Error
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": mpost})
		mylog.Mlogger.Println("Read article successfully")

	})
	return
}

func UpdateArticleAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.PUT("/article/update", func(c *gin.Context) {
		post := sqlmodel.Post{}
		err = c.ShouldBindJSON(&post)
		if err != nil {
			mylog.Mlogger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := db.Where("User_id = ?", post.User_id).Updates(&post)
		if res.Error != nil {
			mylog.Mlogger.Println(res.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			err = res.Error
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Update article successfully"})
		mylog.Mlogger.Println("Update article successfully")
	})
	return
}
func DeleteArticleAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.DELETE("/article", func(c *gin.Context) {
		type Tpost struct {
			Title  string `json:"title,omitempty"`
			Userid uint   `json:"userid,omitempty"`
		}
		dpost := Tpost{}
		err = c.ShouldBindJSON(&dpost)
		if err != nil {
			mylog.Mlogger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mylog.Mlogger.Println("dpost:", dpost)
		//int_userid, _ := strconv.ParseUint(dpost.Userid, 10, 0)
		res := db.Unscoped().Where(sqlmodel.Post{Title: dpost.Title, User_id: dpost.Userid}).Delete(&sqlmodel.Post{})
		if res.Error != nil {
			mylog.Mlogger.Println(res.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
			err = res.Error
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Delete article successfully"})
		mylog.Mlogger.Println("Delete article successfully")
	})
	return
}
