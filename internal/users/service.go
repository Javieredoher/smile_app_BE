package users

import (
	"context"

	"github.com/Javieredoher/smile_app_BE/internal/domain"
)

type UserService interface {
	SaveOdont(ctx context.Context) (domain.Odontologist, error) 
}

type userService struct {
	repository userRepository
}