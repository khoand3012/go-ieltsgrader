package services

import (
	"context"

	"github.com/khoand3012/go-ieltsgrader/internal/domain"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	user, err := us.userRepo.GetByEmail(c, email)
	return user, err
}

// GetUserByEmail(c context.Context, email string) (User, error)
// 	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
// 	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
