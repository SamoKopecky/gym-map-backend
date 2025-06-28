package user

type userPostRequest struct {
	Email string `json:"email"`
}

type userGetRequest struct {
	Id *string `json:"id"`
}
