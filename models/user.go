package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Email   string
	GroupID int
}

func (u *User) GetUser(id int) map[string]interface{} {
	var result = make(map[string]interface{})
	db.Model(&User{}).First(&result, "id = ?", id)
	return result
}
