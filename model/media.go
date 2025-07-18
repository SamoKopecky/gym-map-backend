package model

import "github.com/uptrace/bun"

const YOUTUBE_CONTENT_TYPE = "youtube"

type Media struct {
	bun.BaseModel `bun:"table:media"`
	IdModel
	Timestamp

	Name        string `json:"name"`
	Path        string `json:"path"`
	ContentType string `json:"content_type"`
	UserId      string `json:"user_id"`
}

func NewYoutubeMedia(youtubeLink, userId string) Media {
	return Media{
		Path:        youtubeLink,
		Name:        "youtube",
		ContentType: YOUTUBE_CONTENT_TYPE,
		UserId:      userId,
	}
}
