package apisrc

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myblok/mylog"
	"myblok/sqlmodel"
	"net/http"
	"time"
)

func RegisterAPI(r *gin.Engine, db *gorm.DB) (err error) {
	r.PUT("/register", func(c *gin.Context) {
		user := sqlmodel.User{}
		err = c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			mylog.Mlogger.Println(err)
			return
		}
		hashdPassword, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			err = err2
			mylog.Mlogger.Println(err)
			return
		}
		user.Password = string(hashdPassword)
		err = db.Create(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			mylog.Mlogger.Println(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"messgae": "User registered successfully"})
		mylog.Mlogger.Println("User registered successfully")
	})
	return
}

func LoginAPI(r *gin.Engine, db *gorm.DB) (tokenString string, err error) {
	var user sqlmodel.User
	r.PUT("/login", func(c *gin.Context) {
		err = c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			mylog.Mlogger.Println(err)
			return
		}
		var dbuser sqlmodel.User
		err = db.Where("username = ?", user.Username).First(&dbuser).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			mylog.Mlogger.Println(err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(user.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			mylog.Mlogger.Println(err)
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       dbuser.ID,
			"username": dbuser.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err = token.SignedString(JwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			mylog.Mlogger.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		mylog.Mlogger.Println("Generate token successfully!")
	})
	return
}
