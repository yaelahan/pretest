package services

import (
	"pretest-indihomesmart/entities"
	"pretest-indihomesmart/models"
)

type UserService interface {
	Create(entities.RegisterRequest) models.User
	FindByEmail(string) models.User
}
