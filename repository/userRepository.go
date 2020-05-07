package repository

import (
	"database/sql"
	"fmt"

	"../domain/model"
)

func SelectSql(db *sql.DB, q string, val []interface{}) ([]model.User, error) {

	dataB := []model.User{}
	datasB := model.User{}

	stmt, err := db.Prepare(q)
	if err != nil {
		return dataB, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(val...)
	if err != nil {
		return dataB, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&datasB.ID, &datasB.Name)
		if err != nil {
			return dataB, err
		}
		dataB = append(dataB, datasB)
	}

	return dataB, nil
}

func InsertSqlReturnData(db *sql.DB, q string, val []interface{}) (model.User, error) {

	d := model.User{}
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	if err != nil {
		fmt.Print(err.Error())
		return d, err
	}

	err = stmt.QueryRow(val...).Scan(&d.ID)
	if err != nil {
		fmt.Print("alias", err.Error())
		return d, err
	}

	return d, nil
}
