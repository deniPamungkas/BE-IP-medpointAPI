package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorRequest struct {
	// Id               string `path:"id"`
	UserId           uuid.UUID `json:"user_id"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	SpecializationId uuid.UUID `json:"specialization_id"`
}

type DoctorRequestWithId struct {
	Id               string `path:"id"`
	UserId           uuid.UUID `json:"user_id"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	SpecializationId uuid.UUID `json:"specialization_id"`
}

type DoctorResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Message string
}

type DeleteDoctorResponse struct {
	Success bool `json:"success"`
	Message string
}

type DoctorController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor" type:"custom"`
	Model   models.Doctors
	Payload *DoctorRequest
}

type DoctorControllerWithId struct {
	raiden.ControllerBase
	Http    string `path:"/doctor/{id}" type:"custom"`
	Model   models.Doctors
	Payload *DoctorRequestWithId
}

// get all doctor
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

// get doctor by id
func (c *DoctorControllerWithId) Get(ctx raiden.Context) error {
	doctor := models.Doctors{}
	err := db.NewQuery(ctx).From(models.Doctors{}).Eq("id",c.Payload.Id).Single(&doctor)
	if err != nil {
		return ctx.SendError(err.Error())
	}
	response := DoctorResponse{
		Success: true,
		Data:    doctor,
		Message: "success",
	}
	return ctx.SendJson(response)
}

// add new doctor
func (c *DoctorController) Post(ctx raiden.Context) error {
	payload := models.Doctors{Name: c.Payload.Name, Gender: c.Payload.Gender, UserId: c.Payload.UserId, SpecializationId: c.Payload.SpecializationId}
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
func (c *DoctorControllerWithId) Delete(ctx raiden.Context) error {

	err := db.NewQuery(ctx).From(models.Doctors{}).Eq("id", c.Payload.Id).Delete()

	if err != nil {
		return ctx.SendError("error")
	}

	response := DoctorResponse{
		Success: true,
		Data:    err,
		Message: "success",
	}

	return ctx.SendJson(response)
}

// edit doctor schedule
func (c *DoctorControllerWithId) Patch(ctx raiden.Context) error {

	payload := models.Doctors{Name: c.Payload.Name, Gender: c.Payload.Gender, SpecializationId: c.Payload.SpecializationId, UserId: c.Payload.UserId}

	var updatedDoctor any

	err := db.NewQuery(ctx).From(models.Doctors{}).Eq("id", c.Payload.Id).Update(payload, &updatedDoctor)

	if err != nil {
		return ctx.SendError("error")
	}

	response := DoctorResponse{
		Success: true,
		Data:    updatedDoctor,
		Message: "success update doctor",
	}

	return ctx.SendJson(response)
}