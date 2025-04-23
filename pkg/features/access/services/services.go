package services

import (
	"fmt"
	"merchant-portal/pkg/features/access/models"
	"merchant-portal/pkg/features/access/repositories"
)

type AccessServices struct {
	r repositories.AccessRepository
}

func NewAccessServices(r repositories.AccessRepository) *AccessServices {
	return &AccessServices{
		r: r,
	}
}
func (s *AccessServices) AddPortalUser(user *models.AccessPortalUser) {
	fmt.Println("Adding Portal User:", user)
	s.r.AddPortalUser()
}

func (s *AccessServices) GetPortalUsers() {
	s.r.GetPortalUsers()
	// Add logic to get all portal users
}
func (s *AccessServices) GetPortalUserByID(id string) {
	s.r.GetPortalUserByID(id)
	// Add logic to get a portal user by ID
}
func (s *AccessServices) UpdatePortalUser(id string) {
	s.r.UpdatePortalUser(id)
	// Add logic to update a portal user by ID
}
func (s *AccessServices) DeletePortalUser(id string) {
	s.r.DeletePortalUser(id)
	// Add logic to delete a portal user by ID
}
