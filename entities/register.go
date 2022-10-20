package entities

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,uniquedb=users.email"`
	Phone    string `json:"phone" validate:"required,uniquedb=users.phone"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
