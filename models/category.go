package models

import (
	"time"

	"github.com/google/uuid"
)

// Category - Describe Category data structure
type Category struct {
	ID string           `bson:"ID"`
	Name string         `bson:"name"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"createdAt"`
}

// NewCategory - Create new Category instance with managed ID
func NewCategory(name string) (category Category) {
	id := uuid.New().String()
	createdAt := time.Now()

	return Category{
		ID: id, 
		Name: name, 
		CreatedAt: createdAt, 
		UpdatedAt: createdAt,
	}
}

// CategoryRepo - Describe interface that must be implemented to handle persistence methods for Category
type CategoryRepo interface {
	FindByID(ID string) (category Category, err error)
	Create(category Category) (created Category, err error)
	Delete(ID string) (category Category, err error)
	Override(category Category) (updated Category, err error)
}

// CategoryService - Describe interface that must be implemented to handle persistence methods for Category
type CategoryService interface {
	FindByID(ID string) (category Category, err error)
	Create(name string) (category Category, err error)
}