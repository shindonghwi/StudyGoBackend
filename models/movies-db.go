package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	row := m.DB.QueryRow("select idx, id, title, description, releaseData, runtime, rating, mpaRating, created_at, updated_at from Movie where id = ?", id)

	var movie Movie

	err := row.Scan(
		&movie.IDX,
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		println("Error: ", err.Error())
		return nil, err
	}

	return &movie, nil
}

func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
