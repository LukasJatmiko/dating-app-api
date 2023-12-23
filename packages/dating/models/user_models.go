package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                int            `json:"id" gorm:"column:id;primaryKey"`
	Name              string         `json:"name" gorm:"column:name;" validate:"required"`
	Email             string         `json:"email" gorm:"column:email" validate:"required,email"`
	Password          string         `json:"password" gorm:"column:password" validate:"required"`
	ProfilePictureURL string         `json:"profile_picture_url" gorm:"profile_picture_url"`
	NickName          string         `json:"nickname" gorm:"column:nickname"`
	Birthday          datatypes.Date `json:"birthday" gorm:"column:birthday" validate:"required"`
	GenderID          int            `json:"-" gorm:"column:gender_id"`
	Gender            Gender         `json:"gender" gorm:"foreignKey:gender_id;references:id" validate:"required"`
	GenderInterest    []Gender       `json:"gender_interest" gorm:"many2many:gender_interest;" validate:"required"`
}

type Gender struct {
	gorm.Model
	ID   int    `json:"id" gorm:"column:id;primaryKey;" validate:"required,gte=1"`
	Name string `json:"name" gorm:"gender" validate:"required"`
}

type UserClaims struct {
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	UserEmail string `json:"user_email" gorm:"column:user_email"`
	UserName  string `json:"user_name" gorm:"column:user_name"`
	jwt.RegisteredClaims
}

type Data struct {
	Token string
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
