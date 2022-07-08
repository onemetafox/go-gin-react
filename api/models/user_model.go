package models

import (
	"gorm.io/gorm"
	"github.com/todaycodemaster/go-gin-react/api/database"
)

var db = database.InitDb()

type User struct {
	gorm.Model
	id				int
	first_name		string
	last_name		string
	email			string
	phone			string
	password 		string
	Articles []Article
}

func SaveUser(data map[string]interface{})  (int, error) {
	if data["id"] != nil{
		if err := db.Model(&User{}).Where("id = ?", data["id"].(int)).Updates(data).Error; err != nil {
			return data["id"].(int), nil
		}else{
			return 0, err
		}
	}else{
		user := User{
			first_name: 	data["first_name"].(string),
			last_name:		data["last_name"].(string),
			email: 			data["email"].(string),
			phone: 			data["phone"].(string),
			password: 		data["password"].(string),
		}
		result := db.Create(&user)
		if err := result.Error; err != nil {
			return user.id, nil
		}else{
			return 0, err
		}
	}
	
}
func GetUser(id int) (*User, error) {
	var user User
	err := db.First(&user, 10).Error
	if err != nil{
		return nil, err
	}else{
		return &user, nil
	}
	
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error){
	var users []*User
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func DelUser(id int) error  {
	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}
	return nil
}

