package dto

type RegisterRequestForm struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"first_name" validate:"required,max=100"`
	LastName    string `json:"last_name" validate:"required,max=100"`
	BirthDay    string `json:"birthday" validate:"required,datetime=2006-01-02"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
