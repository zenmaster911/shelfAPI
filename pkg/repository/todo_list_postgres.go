package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/zenmaster911/shelfAPI"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodolistPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	CreateListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(CreateListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("create list error %s", err)
	}

	CreateUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(CreateUsersListQuery, userId, id)

	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("create users list error %s", err)
	}
	return id, tx.Commit()
}
