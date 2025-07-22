package repository

import (
	"database/sql"

	"github.com/JairoCC/go-react/backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
