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

func (service *UserServiceImpl) Create(request entities.RegisterRequest) models.User {
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: utils.HashPassword(request.Password),
	}

	err := service.db.Create(&user).Error
	exceptions.PanicIfNeeded(err)

	return user
}

func (service *UserServiceImpl) FindByEmail(email string) models.User {
	user := models.User{}

	service.db.Where(models.User{Email: email}).Find(&user)

	return user
}
