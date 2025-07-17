package model

import "github.com/uptrace/bun"

type Instruction struct {
	bun.BaseModel `bun:"table:instruction"`
	IdModel
	Timestamp

	UserId      string `json:"user_id"`
	ExerciseId  int    `json:"exercise_id"`
	Description string `json:"description"`
	MediaIds    []int  `json:"media_ids" bun:",array"`
}

func BuildInstruction(userId, description string, exerciseId int, mediaIds []int) Instruction {
	return Instruction{
		UserId:      userId,
		Description: description,
		ExerciseId:  exerciseId,
		MediaIds:    mediaIds,
	}
}
