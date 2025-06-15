package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-shop-diplom/dto"
	"go-shop-diplom/models"
	"go-shop-diplom/services"
	"net/http"
)

func AddToBasket(context *fiber.Ctx) error {
	basketDTO := dto.BasketDTO{}
	err := context.BodyParser(&basketDTO)

	product, err := services.FindProductByName(basketDTO.ProductName)
	if err != nil {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to find product", "error": err.Error()})
		return err
	}

	user, err := services.FindUserByLogin(basketDTO.Login)
	if err != nil {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to find user", "error": err.Error()})
		return err
	}

	ExistedBasket, err := services.FindBasketForUserAndProduct(user, product)
	if err != nil && err.Error() != "record not found" {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to find basket", "error": err.Error()})
		return err
	}

	if err != nil && err.Error() == "record not found" {

		newBasket := models.Basket{
			UserID:    user.ID,
			ProductID: product.ID,
			Count:     basketDTO.Count,
		}

		err = services.CreateBasket(&newBasket)
		if err != nil {
			_ = context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "failed to create basket", "error": err.Error()})
			return err
		} else {
			_ = context.Status(http.StatusOK).JSON(&fiber.Map{
				"message": "Item added to basket successfully",
			})
			return nil
		}
	}

	ExistedBasket.Count += basketDTO.Count
	err = services.UpdateBasket(ExistedBasket)
	if err != nil {
		_ = context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to update basket", "error": err.Error()})
		return err
	}

	_ = context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Basket updated successfully",
	})
	return nil
}

func PayForBasket(context *fiber.Ctx) error {
	basketPayDTO := dto.BasketPayDTO{}
	err := context.BodyParser(&basketPayDTO)
	if err != nil {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "invalid format", "error": err.Error()})
		return err
	}

	user, err := services.FindUserByLogin(basketPayDTO.Login)
	if err != nil {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to find user", "error": err.Error()})
		return err
	}

	baskets, err := services.FindAllBasketsForUser(user)

	if len(baskets) == 0 || (baskets == nil) {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Your basket is empty"})
		return nil
	}
	if err != nil {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Your basket is empty"})
		return err
	}

	if err != nil {
		_ = context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to find basket", "error": err.Error()})
		return err
	}

	sum := int(user.Deposit) - getTotalPrice(&baskets)
	if sum < 0 {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": fmt.Sprintf("There are not enough funds in the balance. Add %d coins", sum)})
		return nil
	}

	user.Deposit = uint32(sum)
	user, err = services.UpdateDeposit(user)
	if err != nil {
		_ = context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to update user deposit", "error": err.Error()})
		return err
	}

	err = services.DeleteAllBasketsForUser(user)
	if err != nil {
		_ = context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to delete baskets", "error": err.Error()})
		return err
	}

	_ = context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": fmt.Sprintf("You have successfully purchased the product. There are %d coins left in your balance", user.Deposit),
	})
	return nil
}

func getTotalPrice(baskets *[]models.Basket) int {
	totalPrice := 0
	for _, basket := range *baskets {
		if basket.Product == nil {
			continue // Skip if product is not found

		}
		totalPrice += int(basket.Product.Cost) * int(basket.Count)
	}
	return totalPrice
}
