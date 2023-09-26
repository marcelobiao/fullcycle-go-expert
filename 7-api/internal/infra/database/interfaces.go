package database

import "github.com/marcelobiao/fullcycle-go-expert/7-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emailid string) (*entity.User, error)
}
