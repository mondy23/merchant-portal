package utils

import "github.com/gofiber/fiber/v2"

func BindBody(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	return nil
}

func FuncHandler(c *fiber.Ctx, function func() (map[string]interface{}, error)) error {
	resp, err := function()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(resp)
}
