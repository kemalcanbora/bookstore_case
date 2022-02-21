package models

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CreatedTime int64  `json:"created_time"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required,max=20,min=1"`
}
