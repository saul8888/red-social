package user

import (
	"github.com/labstack/echo"
	"github.com/orderforme/user/user/model"
)

type Validate interface {
	ValiCreate(context echo.Context) error
	ValiUpdate(context echo.Context) error
	Populate(customer model.User) error
}

type validate struct {
	createUser model.User
	updateUser model.UserUpdate
}

func CreateValidate(data model.User) Validate {
	return &validate{createUser: data}
}

func UpdateValidate(data model.UserUpdate) Validate {
	return &validate{updateUser: data}
}

func (vali *validate) ValiCreate(context echo.Context) error {
	if err := context.Validate(vali.createUser); err != nil {
		return err
	}
	return nil
}

func (vali *validate) ValiUpdate(context echo.Context) error {
	if err := context.Validate(vali.updateUser); err != nil {
		return err
	}
	return nil
}

func (request *validate) Populate(customer model.User) error {
	if request.updateUser.UserName == "" {
		model.Userupdate["userName"] = customer.UserName
	} else {
		model.Userupdate["userName"] = request.updateUser.UserName
	}
	if request.updateUser.FirstName == "" {
		model.Userupdate["firstName"] = customer.FirstName
	} else {
		model.Userupdate["firstName"] = request.updateUser.FirstName
	}

	if request.updateUser.LastName == "" {
		model.Userupdate["lastName"] = customer.LastName
	} else {
		model.Userupdate["lastName"] = request.updateUser.LastName
	}

	if request.updateUser.Email == "" {
		model.Userupdate["email"] = customer.Email
	} else {
		model.Userupdate["email"] = request.updateUser.Email
	}

	if request.updateUser.Password == "" {
		model.Userupdate["password"] = customer.Password
	} else {
		model.Userupdate["password"] = request.updateUser.Password
	}

	if request.updateUser.Avatar == "" {
		model.Userupdate["Avatar"] = customer.Avatar
	} else {
		model.Userupdate["Avatar"] = request.updateUser.Avatar
	}

	if request.updateUser.Front == "" {
		model.Userupdate["front"] = customer.Front
	} else {
		model.Userupdate["front"] = request.updateUser.Front
	}

	if request.updateUser.Biografia == "" {
		model.Userupdate["biografia"] = customer.Biografia
	} else {
		model.Userupdate["biografia"] = request.updateUser.Biografia
	}

	if request.updateUser.Location == "" {
		model.Userupdate["location"] = customer.Location
	} else {
		model.Userupdate["location"] = request.updateUser.Location
	}

	if request.updateUser.Website == "" {
		model.Userupdate["website"] = customer.Website
	} else {
		model.Userupdate["website"] = request.updateUser.Website
	}

	return nil
}
