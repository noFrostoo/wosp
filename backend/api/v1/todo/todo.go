package todo

import (
	"backend/models"

	"github.com/google/uuid"
)

type Store interface {
	GetByID(uuid.UUID) (*models.Todo, error)
	GetAllByUser(string) ([]*models.Todo, error)
	Create(uuid.UUID, string, string, string) error
	Update(*models.Todo) error
	Delete(uuid.UUID)
}