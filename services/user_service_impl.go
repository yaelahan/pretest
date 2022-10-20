package services

import (
	"gorm.io/gorm"
	"pretest-indihomesmart/entities"
	"pretest-indihomesmart/exceptions"
	"pretest-indihomesmart/models"
	"pretest-indihomesmart/utils"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{
		db: db,
	}
}

func (service *UserServiceImpl) Create(request entities.RegisterRequest) entities.RegisterResponse {
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: utils.HashPassword(request.Password),
	}

	err := service.db.Create(&user).Error
	exceptions.PanicIfNeeded(err)

	return entities.RegisterResponse{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}
