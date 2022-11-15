package users

import (
	"context"
	"database/sql"

	"github.com/Javieredoher/smile_app_BE/internal/domain"
)

type UserRepository interface {
	SaveUser(context.Context, domain.Odontologist) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository{
	return &repository{
		db: db,
	}
}

func (r *repository) SaveUser(ctx context.Context, user domain.User) (int, error) {
	query := "INSERT INTO odont(attributes,separated) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.attributes)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}