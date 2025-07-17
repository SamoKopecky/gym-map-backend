package service

import "gym-map/store"

type Media struct {
	MediaCrud store.Media
}

func (m Media) IsTrainerOwned(userId string, id int) (bool, error) {
	media, err := m.MediaCrud.GetById(id)
	if err != nil {
		return false, err
	}

	if media.UserId == userId {
		return true, nil
	}

	return false, nil
}
