package api

import (
	"gym-map/store"
	"net/http"
	"strconv"
)

type NewModulator[M any] interface {
	ToNewModel() M
}

type ExistingModulator[M any] interface {
	ToExistingModel(id int) M
}

func DeleteModel[M any](cc *DbContext, crud store.StoreBase[M]) error {
	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	err = crud.Delete(id)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func PatchModel[R ExistingModulator[M], M any](cc *DbContext, crud store.StoreBase[M]) error {
	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	params, err := BindParams[R](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.ToExistingModel(id)
	err = crud.Update(&model)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func PostModel[R NewModulator[M], M any](cc *DbContext, crud store.StoreBase[M]) error {
	params, err := BindParams[R](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newModel := params.ToNewModel()
	err = crud.Insert(&newModel)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, newModel)
}

func GetModels[M any](cc *DbContext, crud store.StoreBase[M]) error {
	models, err := crud.Get()
	if err != nil {
		return err
	}

	if models == nil {
		models = []M{}
	}

	return cc.JSON(http.StatusOK, models)
}
