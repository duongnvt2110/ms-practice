package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthProfile struct {
	Id        int       `gorm:"column:id" json:"id"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}


type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Custom claims for JWT
type AuthClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}