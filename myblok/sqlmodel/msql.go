package sqlmodel

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primarykey" ;foreignkey:"User_id"`
	Username string
	Password string
	Email    string
}

type Post struct {
	gorm.Model
	ID      uint `gorm:"primarykey" ;foreignkey:"Post_id"`
	Title   string
	Content string
	User_id uint
	User    User
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Title == "" {
		return errors.New("Missing title")
	} else if p.Content == "" {
		return errors.New("Missing content")
	}
	return nil
}

func (p *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	// select title and user is true
	res := tx.Where(&Post{Title: p.Title, User_id: p.User_id}).First(&Post{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *Post) BeforeDelete(tx *gorm.DB) (err error) {
	// select title and user is true
	res := tx.Where(&Post{Title: p.Title, User_id: p.User_id}).First(&Post{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

type Comment struct {
	gorm.Model
	ID      uint `gorm:"primarykey"`
	Content string
	User_id uint
	Post_id uint
	Post    Post
	User    User
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Post{}, &Comment{}, &User{})
}
