package instruction

import (
	"gym-map/api"
	"gym-map/model"
)

type instructionGetRequest struct {
	ExerciseId *int    `query:"exercise_id"`
	UserId     *string `query:"string_id"`
}

type instructionPostRequest struct {
	UserId      string `json:"user_id"`
	ExerciseId  int    `json:"exercise_id"`
	Description string `json:"description"`
}

type instructionMediaPostRequest struct {
	YoutubeVideoId *string `json:"youtube_video_id"`
	Name           *string `json:"name"`
}

func (ipr instructionPostRequest) ToNewModel() model.Instruction {
	return model.BuildInstruction(ipr.UserId, ipr.Description, ipr.ExerciseId, []int{})
}

type instructionPatchRequest struct {
	Description *string `json:"description"`
}

func (ipr instructionPatchRequest) ToExistingModel(id int) model.Instruction {
	return model.Instruction{
		IdModel:     model.IdModel{Id: id},
		Description: api.DerefString(ipr.Description),
	}
}

type instructionPostResponse struct {
	MediaIds []int `json:"media_ids"`
}
