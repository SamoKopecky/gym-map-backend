package user

type userPostRequest struct {
	Email string `json:"email"`
}

type userPatchResponse struct {
	AvatarId string `json:"avatar_id"`
}
