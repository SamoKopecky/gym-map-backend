package model

type User struct {
	Id string `json:"id"`
	UserBase
}

type UserBase struct {
	Email     string  `json:"email"`
	Name      *string `json:"name"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
