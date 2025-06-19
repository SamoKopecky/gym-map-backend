package model

import "github.com/uptrace/bun"

type Instruction struct {
	bun.BaseModel `bun:"table:instruction"`
	IdModel
	Timestamp

	UserId      string  `json:"user_id"`
	ExerciseId  int     `json:"exercise_id"`
	Description string  `json:"description"`
	FileId      *string `json:"file_id"`
	FileName    *string `json:"file_name"`
}

func BuildInstruction(userId, description string, exerciseId int, fileId, fileName *string) Instruction {
	return Instruction{
		UserId:      userId,
		Description: description,
		ExerciseId:  exerciseId,
		FileId:      fileId,
		FileName:    fileName,
	}
}
