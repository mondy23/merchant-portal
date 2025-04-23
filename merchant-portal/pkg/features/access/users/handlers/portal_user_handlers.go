package handlers

import "github.com/gofiber/fiber/v2"

// ACCESS

// GetPortalUsers handles the GET /portal-users endpoint (list all users)
func GetPortalUsers(c *fiber.Ctx) error {
	return c.SendString("Get All Users Logic")
}

// AddPortalUser handles the POST /portal-users endpoint (create a new user)
func AddPortalUser(c *fiber.Ctx) error {
	return c.SendString("Add User Logic")
}

// GetPortalUser handles the GET /portal-users/:id endpoint (get user by ID)
func GetPortalUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get User Logic for ID: " + id)
}

// UpdatePortalUser handles the PUT /portal-users/:id endpoint (update user by ID)
func UpdatePortalUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update User Logic for ID: " + id)
}

// DeletePortalUser handles the DELETE /portal-users/:id endpoint (delete user by ID)
func DeletePortalUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete User Logic for ID: " + id)
}
