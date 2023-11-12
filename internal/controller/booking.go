package controller

import (
	"fww-wrapper/internal/data/dto"
	"fww-wrapper/internal/data/dto_booking"
	"fww-wrapper/internal/tools"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Booking(ctx *fiber.Ctx) error {
	var req dto_booking.Request
	if err := ctx.BodyParser(&req); err != nil {
		log.Println(err)
		c.Log.Error(err)
		response := tools.ResponseBuilder(nil, dto.MetaResponse{
			StatusCode: "ERR400",
			IsSuccess:  false,
			Message:    "Bad Request",
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := c.Adapter.Booking(&req)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) GetDetailBooking(ctx *fiber.Ctx) error {
	codeBooking := ctx.Query("code_booking", "")

	result, err := c.Adapter.GetDetailBooking(codeBooking)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	meta := dto.MetaResponse{
		StatusCode: "200",
		IsSuccess:  true,
		Message:    "Success",
	}

	response := tools.ResponseBuilder(result, meta)

	return ctx.Status(fiber.StatusOK).JSON(response)
}