package user

import (
	"errors"
	"virtual-white-board-service/internal/dto"
	"virtual-white-board-service/internal/models"
	"virtual-white-board-service/internal/services"

	"gorm.io/gorm"
)

//Login interface for user login
type Login interface {
	Login(user *models.User) (*dto.LoginResponse, error)
}

type login struct {
	db         *gorm.DB
	jwtService services.JWTService
}

//NewLogin creates new login validator
func NewLogin(db *gorm.DB, jwtService services.JWTService) Login {
	res := login{db: db, jwtService: jwtService}
	return res
}

//Insert persists the post to the database
func (m login) Login(user *models.User) (*dto.LoginResponse, error) {
	var dbUser *models.User
	m.db.First(&dbUser, "username = ?", user.Username)
	if dbUser != nil && dbUser.Password == user.Password {
		token := m.jwtService.GenerateToken(user.Username)
		return &dto.LoginResponse{Username: user.Username, Token: token}, nil

	}
	return nil, errors.New("Invalid login")
}
