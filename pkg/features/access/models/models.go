package models

import "golang.org/x/crypto/bcrypt"

// BaseAccessPortalUser contains common fields for user representations.
type BaseAccessPortalUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleInitial string `json:"middle_initial"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// AccessPortalUserInput represents the data a user sends for registration or login.
type AccessPortalUserInput struct {
	Username string `json:"username" binding:"required"` // Add validation tags
	Password string `json:"password" binding:"required"`
}

// AccessPortalUser represents the internal representation of a portal user.
type AccessPortalUser struct {
	BaseAccessPortalUser
	PasswordHash string `json:"-"` // Store the bcrypt hash, omit from JSON
}

// AccessPortalUserResponse represents the data sent in API responses (without sensitive info).
type AccessPortalUserResponse struct {
	BaseAccessPortalUser
}

// SetPassword hashes the provided password and stores it in the PasswordHash field.
func (u *AccessPortalUser) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

// ComparePassword compares the provided password with the stored hash.
func (u *AccessPortalUser) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

// NewAccessPortalUserResponse creates a safe response object from the internal user model.
func NewAccessPortalUserResponse(user *AccessPortalUser) *AccessPortalUserResponse {
	return &AccessPortalUserResponse{
		BaseAccessPortalUser: BaseAccessPortalUser{
			ID:            user.ID,
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			MiddleInitial: user.MiddleInitial,
			Name:          generateFullName(user.FirstName, user.MiddleInitial, user.LastName),
			Role:          user.Role,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
		},
	}
}

// generateFullName constructs the full name.
func generateFullName(first, middle, last string) string {
	fullName := first
	if middle != "" {
		fullName += " " + middle
	}
	fullName += " " + last
	return fullName
}

// BeforeSave hook to automatically generate the Name field before saving to the database.
func (u *AccessPortalUser) BeforeSave() {
	u.Name = generateFullName(u.FirstName, u.MiddleInitial, u.LastName)
}
