package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinit-chauhan/go-bloomservice/internal/bloom"
)

func AddHandler(c *fiber.Ctx) error {
	// This handler will handle the addition of items to the bloom filter.
	var request struct {
		Item string `json:"item" validate:"required"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Item == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Item is required",
		})
	}

	// Add the item to the bloom filter
	bloom.Filter.Add(request.Item)

	// Return a success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Item added to bloom filter successfully",
		"item":    request.Item,
	})
}

func CheckHandler(c *fiber.Ctx) error {
	// This handler will handle the addition of items to the bloom filter.
	var request struct {
		Item string `json:"item" validate:"required"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Item == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Item is required",
		})
	}

	// Check if the item exists in the bloom filter
	exists := bloom.Filter.Exists(request.Item)
	if exists {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"exists": true,
			"item":   request.Item,
		})
	}

	// If the item does not exist, return a 404 Not Found response
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"exists": false,
		"item":   request.Item,
	})

}

func StatsHandler(c *fiber.Ctx) error {
	// This handler will handle the addition of items to the bloom filter.
	params := bloom.Filter.GetParameters()
	stats := bloom.Filter.GetStatistics()

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{"params": params, "stats": stats})
}

func ResetHandler(c *fiber.Ctx) error {
	// This handler will handle the addition of items to the bloom filter.
	bloom.Filter.Clear()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Bloom filter reset successfully",
	})
}
