package model

import (
	"errors"

	"github.com/google/uuid"
	"github.com/michee/authentificationApi/database"
	"github.com/michee/authentificationApi/utils"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	UserId   string `gorm:"primary_key;column:userId"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	Token    string `json:"token"`
}


func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.UserId = uuid.New().String()
	return nil
}

func init() {
	database.ConnectDB()
	DB = database.GetDB()
	// DB.DropTableIfExists(&User{})
	DB.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User{
	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword
	DB.Create(u)
	return u
}


func GetAllUser() []User{
	var u []User
	DB.Find(&u)
	return u
}


func GetUserById (Id string) (*User, *gorm.DB) {
	var getUser User
	db := DB.Where("userId=?", Id).Find(&getUser)

	return &getUser, db
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func DeleteUser(Id string) User{
	var u User
	DB.Where("userId=?", Id).Delete(&u)
	return u
}