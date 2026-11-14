package notes

import (
	"database/sql"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(body string) (int64, error) {
	res, err := s.db.Exec("INSERT INTO notes (title, body) VALUES (?,?)", "Sem título", body)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
