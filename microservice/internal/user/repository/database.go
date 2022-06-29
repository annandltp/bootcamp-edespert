package repository

import (
	"context"
	"course/internal/domain"

	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{
		db: db,
	}
}

func (dbr DBRepository) GetByID(ctx context.Context, userID int) (domain.User, error) {
	var user domain.User
	err := dbr.db.WithContext(ctx).Where("id = ?", userID).Take(&user).Error
	return user, err
}

func (dbr DBRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := dbr.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	return user, err
}
