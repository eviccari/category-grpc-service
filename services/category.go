package services

import "github.com/eviccari/category-grpc-service/models"

// CategoryService - Describe CategoryService structure
type CategoryService struct {
	repo models.CategoryRepo 
}

// FindByID - Find category searching by ID
func (cs *CategoryService) FindByID(ID string) (category models.Category, err error){
	c, err := cs.repo.FindByID(ID)
	if err != nil {
		return models.Category{}, err
	}

	return c, nil 
}

// Create - Create new category into repository
func (cs *CategoryService) Create(name string) (category models.Category, err error){
	c := models.NewCategory(name)
	created, err := cs.repo.Create(c)
	if err != nil {
		return models.Category{}, err
	}

	return created, nil
}


