package user

import (
	"context"
	"course/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserRepo interface {
	GetByID(ctx context.Context, userID int) (domain.User, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

var signatureKey = []byte("mySuperSecretSignature")

func (us UserService) DecriptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return signatureKey, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}
	if !parsedToken.Valid {
		return data, errors.New("token invalid")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func (us UserService) IsUserExists(ctx context.Context, userID int) bool {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	user, err := us.repo.GetByID(ctx, userID)
	return err == nil && user.ID > 0
}
