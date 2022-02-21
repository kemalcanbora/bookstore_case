package models

type JwtUserAuth struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	Exp        string `json:"exp"`
	FirstName  string `json:"first_name"`
}
