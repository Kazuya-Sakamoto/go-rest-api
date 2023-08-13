package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IDiaryRepository interface {
	GetAllDiaries(diaries *[]model.Diary, userId uint) error
}

type diaryRepository struct {
	db *gorm.DB
}

func NewDiaryRepository(db *gorm.DB) IDiaryRepository {
	return &diaryRepository{db}
}

func (dr *diaryRepository) GetAllDiaries(diaries *[]model.Diary, userId uint) error {
	if err := dr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(diaries).Error; err != nil {
		return err
	}
	return nil
}
