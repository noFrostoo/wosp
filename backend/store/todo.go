package store

import (
	"backend/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TodoStore struct {
	db *sqlx.DB
}

func NewTodoStore(db *sqlx.DB) *TodoStore {
	return &TodoStore{
		db: db,
	}
}

func (ts TodoStore) GetByID(id uuid.UUID) (*models.Todo, error) {
	t := &models.Todo{}
	if err := ts.db.Get(t, "SELECT * FROM \"todo\" where id=$1", id); err != nil {
		return nil, err
	}

	return t, nil
}

func (ts TodoStore) GetAllByUser(user_id uuid.UUID) (*[]models.Todo, error) {
	t := &[]models.Todo{}
	if err := ts.db.Select(t, "SELECT * FROM \"todo\" where user_id=$1", user_id); err != nil {
		return nil, err
	}

	return t, nil
}

func (ts TodoStore) Create(t *models.Todo) (*models.Todo, error) {
	if err := ts.db.Get(t, "Insert into \"todo\" (user_id, title, description, done, due_at) values ($1, $2, $3, $4, $5) returning *", t.User_id, t.Title, t.Description, t.Done, t.Due_at); err != nil {
		return nil, err
	}

	return t, nil

}

func (ts TodoStore) Update(t *models.Todo) (*models.Todo, error) {
	if _, err := ts.db.Exec("update \"todo\" set user_id=$2, title=$3, description=$4, done=$5, due_at=$5 where id = $1", t.Id, t.User_id, t.Title, t.Description, t.Done, t.Due_at); err != nil {
		return nil, err
	}

	new_t := &models.Todo{}
	if err := ts.db.Get(new_t, "SELECT * FROM \"todo\" where id=$1", t.Id); err != nil {
		return nil, err
	}

	return new_t, nil
}

func (ts TodoStore) Delete(id uuid.UUID) error {
	if _, err := ts.db.Exec("delete from \"todo\" where id = $1", id); err != nil {
		return err
	}

	return nil
}
