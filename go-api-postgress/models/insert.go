package models

import (
	"api/database"
)

func Insert(todo Todo) (id int64, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return 
	}

	defer conn.Close()

	transaction := `INSERT INTO todos (Title, Description, Done) VALUES ($1, $2, $3) RETURNING Id`

	err = conn.QueryRow(transaction, todo.Title, todo.Description, todo.Done).Scan(&id)

	return 
}