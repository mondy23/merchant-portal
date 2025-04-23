package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type AccessRepository interface {
	AddPortalUser()
	GetPortalUsers()
	GetPortalUserByID(id string)
	UpdatePortalUser(id string)
	DeletePortalUser(id string)
}

type Database struct {
	db *pgxpool.Pool
}

func NewAccessRepository(db *pgxpool.Pool) *Database {
	return &Database{
		db: db,
	}
}

func (db *Database) AddPortalUser() {
	// Add logic to add a portal user
}
func (db *Database) GetPortalUsers() {
	// Add logic to get all portal users
}
func (db *Database) GetPortalUserByID(id string) {
	// Add logic to get a portal user by ID
}
func (db *Database) UpdatePortalUser(id string) {
	// Add logic to update a portal user by ID
}
func (db *Database) DeletePortalUser(id string) {
	// Add logic to delete a portal user by ID
}
