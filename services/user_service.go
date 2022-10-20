package services

import (
	"pretest-indihomesmart/entities"
)

type UserService interface {
	Create(request entities.RegisterRequest) entities.RegisterResponse
}
