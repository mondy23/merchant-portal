package handlers

import (
	"merchant-portal/pkg/features/access/models"
	"merchant-portal/pkg/features/access/services"
	utils "merchant-portal/pkg/utils/handler"

	"github.com/gofiber/fiber/v2"
)

// ACCESS/PORTAL USER HANDLER
type AccessHandlers struct {
	s *services.AccessServices
}

func NewAccessHandlers(s *services.AccessServices) *AccessHandlers {
	return &AccessHandlers{
		s: s,
	}
}

func (h *AccessHandlers) GetPortalUsers(c *fiber.Ctx) error {
	return c.SendString("Get All Users Logic")
}

func (h *AccessHandlers) AddPortalUser(c *fiber.Ctx) error {
	var user models.AccessPortalUser
	utils.BindBody(c, &user)
	h.s.AddPortalUser(&user)
	return c.JSON(user)
}

func (h *AccessHandlers) GetPortalUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get User Logic for ID: " + id)
}

func (h *AccessHandlers) UpdatePortalUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update User Logic for ID: " + id)
}

func (h *AccessHandlers) DeletePortalUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete User Logic for ID: " + id)
}
