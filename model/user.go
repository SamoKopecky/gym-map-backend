package model

type User struct {
	Id       string  `json:"id"`
	Email    string  `json:"email"`
	AvatarId *string `json:"avatar_id"`
	UserBase
}

type UserBase struct {
	Name      *string `json:"name"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
