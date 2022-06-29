package repository

import "gorm.io/gorm"

type ExerciseRepo struct {
	db *gorm.DB
}
