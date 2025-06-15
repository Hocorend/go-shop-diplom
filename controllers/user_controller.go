package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-shop-diplom/dto"
	"go-shop-diplom/services"
)

func AddDeposit(context *fiber.Ctx) error {
	depositDTO := dto.DepositDTO{}
	err := context.BodyParser(&depositDTO)
	login := depositDTO.Login
	amount := depositDTO.Amount

	user, err := services.FindUserByLogin(login)
	if err != nil {
		_ = context.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to find user", "error": err.Error()})
		return nil
	}

	user.Deposit += amount
	user, err = services.UpdateDeposit(user)
	if err != nil {
		return fmt.Errorf("failed to add deposit: %w", err)
	}

	_ = context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": fmt.Sprintf("The deposit was successfully replenished to %d, the current deposit is %d", amount, user.Deposit),
	})
	return nil
}
