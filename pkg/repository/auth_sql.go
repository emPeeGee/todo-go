package repository

import (
	"fmt"

	"github.com/emPeeGee/todo-go"
	"github.com/jmoiron/sqlx"
)

type AuthSql struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (r *AuthSql) CreateUser(user todo.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}
