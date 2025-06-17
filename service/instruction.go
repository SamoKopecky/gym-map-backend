package service

import (
	"gym-map/fetcher"
	"gym-map/model"
	"gym-map/schema"
	"gym-map/store"
)

type Instruction struct {
	IAM             fetcher.IAM
	InstructionCrud store.Instruction
}

func (i Instruction) IsTrainerOwned(userId string, id int) (bool, error) {
	instruction, err := i.InstructionCrud.GetById(id)
	if err != nil {
		return false, err
	}

	if instruction.UserId == userId {
		return true, nil
	}

	return false, nil
}

func (i Instruction) Get() (userInstructions []schema.Instruction, err error) {
	instructions, err := i.InstructionCrud.Get()
	if err != nil {
		return
	}
	userInstructions, err = i.withUsers(instructions)
	if err != nil {
		return
	}
	return
}

func (i Instruction) GetByExerciseId(exerciseId int) (userInstructions []schema.Instruction, err error) {
	instructions, err := i.InstructionCrud.GetByExerciseId(exerciseId)
	if err != nil {
		return
	}
	userInstructions, err = i.withUsers(instructions)
	if err != nil {
		return
	}
	return
}

func (i Instruction) GetByUserId(userId string) (userInstructions []schema.Instruction, err error) {
	instructions, err := i.InstructionCrud.GetByUserId(userId)
	if err != nil {
		return
	}

	// TODO: Use single query for single user id
	userInstructions, err = i.withUsers(instructions)
	if err != nil {
		return
	}
	return
}

func (i Instruction) Insert(instruction *model.Instruction) (userInstruction schema.Instruction, err error) {
	err = i.InstructionCrud.Insert(instruction)
	if err != nil {
		return
	}

	// TODO: Use single query for single user id
	userInstructions, err := i.withUsers([]model.Instruction{*instruction})
	if err != nil {
		return
	}

	if len(userInstructions) < 1 {
		return
	}

	userInstruction = userInstructions[0]
	return

}

func (i Instruction) withUsers(instructions []model.Instruction) (userInstructions []schema.Instruction, err error) {
	// TODO: Cache fetching users
	users, err := i.IAM.GetUsers()
	if err != nil {
		return
	}
	usersMap := make(map[string]fetcher.KeycloakUser)
	for _, user := range users {
		usersMap[user.Id] = user
	}

	for _, instruction := range instructions {
		if user, ok := usersMap[instruction.UserId]; ok {
			userInstructions = append(userInstructions, schema.Instruction{
				Instruction: instruction,
				User:        user.ToUserModel(),
			})
		}
	}
	return

}
