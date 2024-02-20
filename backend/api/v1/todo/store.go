package todo

import (
	"backend/models"

	"github.com/google/uuid"
)

type Store interface {
	GetByID(uuid.UUID) (*models.Todo, error)
	GetAllByUser(uuid.UUID) (*[]models.Todo, error)
	Create(*models.Todo) (*models.Todo, error)
	Update(*models.Todo) (*models.Todo, error)
	Delete(uuid.UUID) error
}
