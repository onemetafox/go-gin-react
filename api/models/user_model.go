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

func SaveUser(User *User)  int {
	if User.id != 0{
		result := db.Save(User)
		if result.Error != nil{
			return 0
		}
		return User.id
	}else{
		result := db.Create(User)
		if result.Error != nil{
			return 0
		}
		return User.id
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

