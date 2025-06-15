package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-shop-diplom/models"
	"go-shop-diplom/services"
	"net/http"
)

func GetProducts(context *fiber.Ctx) error {
	products, err := services.FindAllProducts()

	if err != nil {
		_ = context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to get products", "error": err.Error()})
		return err
	}

	_ = context.Status(http.StatusOK).JSON(&fiber.Map{
		"products": models.MapProductsToProductDTOs(products),
	})
	return nil
}
