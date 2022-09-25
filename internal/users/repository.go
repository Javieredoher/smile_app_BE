package users

import (
	"context"
	"database/sql"

	"github.com/Javieredoher/smile_app_BE/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveOdont(context.Context, domain.Odontologist) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return: &repository {
		db: db,
	}
}

func (r *repository) SaveOdont(ctx context.Context, odo domain.Odontologist) (int, error) {
	
	err := db.Create(&odo)
}