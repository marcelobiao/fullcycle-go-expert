package database

import "github.com/marcelobiao/fullcycle-go-expert/7-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emailid string) (*entity.User, error)
}

type ProductInterface interface {
	Create(user *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
