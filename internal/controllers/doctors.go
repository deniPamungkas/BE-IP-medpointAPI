package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorRequest struct {
	Id               string `path:"id"`
	UserId           uuid.UUID
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	SpecializationId *uuid.UUID
}

type DoctorResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Message string
}

type DoctorController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor" type:"custom"`
	Model   models.Doctors
	Payload *DoctorRequest
}

type DeleteDoctorController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor/{id}" type:"custom"`
	Model   models.Doctors
	Payload *DoctorRequest
}

// get all doctor schedule
func (c *DoctorController) Get(ctx raiden.Context) error {
	var doctors []models.Doctors
	err := db.NewQuery(ctx).From(models.Doctors{}).Get(&doctors)
	if err != nil {
		return ctx.SendError(err.Error())
	}
	response := DoctorResponse{
		Success: true,
		Data:    doctors,
		Message: "success",
	}
	return ctx.SendJson(response)
}

// add new doctor
func (c *DoctorController) Post(ctx raiden.Context) error {
	payload := models.Doctors{Name: c.Payload.Name, Gender: c.Payload.Gender}
	var insertedDoctor any
	err := db.NewQuery(ctx).From(models.Doctors{}).Insert(&payload, &insertedDoctor)
	if err != nil {
		return ctx.SendError(err.Error())
	}
	response := DoctorResponse{
		Success: true,
		Data:    insertedDoctor,
		Message: "success add new doctor",
	}
	return ctx.SendJson(response)
}

// delete doctor schedule
func (c *DeleteDoctorController) Delete(ctx raiden.Context) error {

	err := db.NewQuery(ctx).From(models.Doctors{}).Eq("id", c.Payload.Id).Delete()

	if err != nil {
		return ctx.SendError("error")
	}

	response := DoctorResponse{
		Success: true,
		Data:    c.Payload.Id,
		Message: "success",
	}

	return ctx.SendJson(response)
}
