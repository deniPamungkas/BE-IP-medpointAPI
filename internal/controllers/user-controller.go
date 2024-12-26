package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type UserRequest struct{
	Id string `path:"id"`
}

type UserResponse struct{
	Success bool `json:"success"`
	Data any `json:"data"`
}

type UserController struct{
	raiden.ControllerBase
	Http string `path:"/user/{id}" type:"custom"`
	Model models.PublicUsers
	Payload *UserRequest
}

func(c *UserController) Get(ctx raiden.Context) error{
	user := models.PublicUsers{}

	err:=db.NewQuery((ctx)).From(models.PublicUsers{}).Eq("id", c.Payload.Id).Single(&user)

	if err != nil {
		return ctx.SendError(err.Error())
	}

	response:=UserResponse{
		Success: true,
		Data: user,
	}

	return ctx.SendJson(response)
}