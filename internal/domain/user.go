package domain

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const TableUser = "users"

type User struct {
	gorm.Model
	FirstName     string `gorm:"type:varchar(50);" json:"first_name"`
	LastName      string `gorm:"type:varchar(50);" json:"last_name"`
	Email         string `gorm:"type:varchar(100);unique;" json:"email"`
	Salt          string `gorm:"not null" json:"salt"`
	HasedPassword string `gorm:"column:hashed_password;not null" json:"-"`
}

func (u User) TableName() string {
	return TableUser
}

func NewUser(email, password, firstName, lastName string) *User {
	salt := fmt.Sprintf("%d", time.Now().UnixMilli())

	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		Salt:          salt,
		HasedPassword: password + salt,
	}
}

func (u User) VerifyPassword(password string) bool {
	return u.HasedPassword == password+u.Salt
}
