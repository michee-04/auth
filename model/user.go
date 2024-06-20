package model

import (
	"errors"

	"github.com/google/uuid"
	"github.com/michee/authentificationApi/database"
	"github.com/michee/authentificationApi/utils"

	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	UserId   string `gorm:"primary_key;column:user_id"`
	Name     string `gorm:"column:name" json:"name"`
	UserName string `gorm:"column:username" json:"username"`
	Email    string `gorm:"unique;column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Token    string `gorm:"column:token" json:"token"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UserId = uuid.New().String()
	return
}

func init() {
	database.ConnectDB()
	DB = database.GetDB()
	// DB.Migrator().DropTable(&User{})
	if DB != nil {
		err := DB.AutoMigrate(&User{})
		if err != nil {
			panic("Failed to migrate User model: " + err.Error())
		}
	} else {
		panic("DB connection is nil")
	}
}

func (u *User) CreateUser() *User {
	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword
	DB.Create(u)
	return u
}

func GetAllUser() []User {
	var users []User
	DB.Find(&users)
	return users
}

func GetUserById(id string) (*User, *gorm.DB) {
	var user User
	db := DB.Where("user_id = ?", id).First(&user)
	return &user, db
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func DeleteUser(id string) User {
	var user User
	DB.Where("user_id = ?", id).Delete(&user)
	return user
}

func (u *User) Logout() error {
	u.Token = "" // Effacer le token enregistré dans l'objet User
	return DB.Save(&u).Error
}
