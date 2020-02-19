package skeleton

import (
	"context"

	userEntity "go-tutorial-2020/internal/entity/user"
)

// UserData ...
type UserData interface {
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
}

// Service ...
type Service struct {
	userData UserData
}

// New ...
func New(userData UserData) Service {
	return Service{
		userData: userData,
	}
}

// GetAllUsers ...
func (s Service) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
	users, err := s.userData.GetAllUsers(ctx)
	return users, err
}
