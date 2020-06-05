package follow

import (
	"github.com/labstack/echo"
	"github.com/orderforme/user/public/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.Public) error
}

type validate struct {
	createPublic model.Public
	updatePublic model.PublicUpdate
}

func CreateValidate(data model.Public) Validate {
	return &validate{createPublic: data}
}

func UpdateValidate(data model.PublicUpdate) Validate {
	return &validate{updatePublic: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createPublic); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updatePublic); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.Public) error {
	/*
		if request.updatePublic.UserID == "" {
			model.Publicupdate["userId"] = customer.UserID
		} else {
			model.Publicupdate["userId"] = request.updatePublic.UserID
		}
	*/
	if request.updatePublic.Message == "" {
		model.Publicupdate["message"] = customer.Message
	} else {
		model.Publicupdate["message"] = request.updatePublic.Message
	}

	return nil
}
