package model

import "github.com/uptrace/bun"

type Instruction struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp

	UserId      string `json:"user_id"`
	ExerciseId  int    `json:"exercise_id"`
	Description string `json:"description"`
}

func BuildInstruction(userId, description string, exerciseId int) Instruction {
	return Instruction{
		UserId:      userId,
		Description: description,
		ExerciseId:  exerciseId,
	}
}
