package models

import "api/database"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET Title=$2, Description=$3, Done=$4 WHERE Id=$1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}