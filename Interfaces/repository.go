package interfaces

import "sayhello/pkg/models"

type Repository interface {
	Save(item *models.Item) error

	Get(id string) (*models.Item, error)

	GetAll() (models.Items, error)

	Delete(id string) error
}
