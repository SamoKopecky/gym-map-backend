package model

import "github.com/uptrace/bun"

type Media struct {
	bun.BaseModel `bun:"table:media"`
	IdModel
	Timestamp

	OriginalFileName string `json:"original_file_name"`
	DiskFileName     string `json:"disk_file_name"`
	ContentType      string `json:"content_type"`
}
