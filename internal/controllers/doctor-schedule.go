package controllers

import (
	"ipmedpointsistem/internal/models"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type ScheduleRequest struct {
	Id           string `path:"id"`
	DoctorId     *uuid.UUID
	AvailableDay string `json:"available_day"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}

type DeleteScheduleRequest struct {
	Id string `path:"id"`
}

type ScheduleResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Message string
}

type ScheduleController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor-schedule" type:"custom"`
	Model   models.DoctorSchedule
	Payload *ScheduleRequest
}

type DeleteScheduleController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor-schedule/{id}" type:"custom"`
	Model   models.DoctorSchedule
	Payload *DeleteScheduleRequest
}

type UpdateScheduleController struct {
	raiden.ControllerBase
	Http    string `path:"/doctor-schedule/{id}" type:"custom"`
	Model   models.DoctorSchedule
	Payload *ScheduleRequest
}

// get all doctor schedule
func (c *ScheduleController) Get(ctx raiden.Context) error {
	var schedules []models.DoctorSchedule

	err := db.NewQuery(ctx).From(models.DoctorSchedule{}).Get(&schedules)

	if err != nil {
		return ctx.SendError(err.Error())
	}

	response := ScheduleResponse{
		Success: true,
		Data:    schedules,
		Message: "success",
	}
	return ctx.SendJson(response)
}

// add new doctor schedule
func (c *ScheduleController) Post(ctx raiden.Context) error {
	payload := models.DoctorSchedule{AvailableDay: c.Payload.AvailableDay, StartTime: c.Payload.StartTime, EndTime: c.Payload.EndTime}
	var insertedDoctorSchedule any
	err := db.NewQuery(ctx).From(models.DoctorSchedule{}).Insert(&payload, &insertedDoctorSchedule)
	if err != nil {
		return ctx.SendError(err.Error())
	}
	response := DoctorResponse{
		Success: true,
		Data:    insertedDoctorSchedule,
		Message: "success",
	}
	return ctx.SendJson(response)
}

// delete doctor schedule
func (c *DeleteScheduleController) Delete(ctx raiden.Context) error {

	err := db.NewQuery(ctx).From(models.DoctorSchedule{}).Eq("id", c.Payload.Id).Delete()

	if err != nil {
		return ctx.SendError("error")
	}

	response := ScheduleResponse{
		Success: true,
		Data:    c.Payload.Id,
		Message: "success",
	}

	return ctx.SendJson(response)
}

// edit doctor schedule
func (c *UpdateScheduleController) Put(ctx raiden.Context) error {

	payload := models.DoctorSchedule{AvailableDay: c.Payload.AvailableDay, StartTime: c.Payload.StartTime, EndTime: c.Payload.EndTime}

	var updatedSchedule any

	err := db.NewQuery(ctx).From(models.DoctorSchedule{}).Eq("id", c.Payload.Id).Update(payload, &updatedSchedule)

	if err != nil {
		return ctx.SendError("error")
	}

	response := DoctorResponse{
		Success: true,
		Data:    updatedSchedule,
	}

	return ctx.SendJson(response)
}
