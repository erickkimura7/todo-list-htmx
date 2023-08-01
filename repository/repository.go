package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

type Database struct {
	db *sql.DB
}

type DatabaseImpl interface {
	GetAllTodos() ([]Todo, error)
	DeleteById(id string) error
	UpdateStatus(id string) error
}

func New() (*Database, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) GetAllTodos() ([]Todo, error) {
	d.db
}

func (d *Database) DeleteById() error {

}

func (d *Database) UpdateStatus() error {

}
