package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDiaryRepository interface {
	GetAllDiaries(diaries *[]model.Diary, userId uint) error
	GetDiaryById(diary *model.Diary, userId uint, diaryId uint) error
	CreateDiary(diary *model.Diary) error
	UpdateDiary(diary *model.Diary, userId uint, diaryId uint) error
	DeleteDiary(userId uint, diaryId uint) error
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

func (dr *diaryRepository) GetDiaryById(diary *model.Diary, userId uint, diaryId uint) error {
	if err := dr.db.Joins("User").Where("user_id=?", userId).First(diary, diaryId).Error; err != nil {
		return err
	}
	return nil
}

func (dr *diaryRepository) CreateDiary(diary *model.Diary) error {
	if err := dr.db.Create(diary).Error; err != nil {
		return err
	}
	return nil
}

func (dr *diaryRepository) UpdateDiary(diary *model.Diary, userId uint, diaryId uint) error {
	result := dr.db.Model(diary).Clauses(clause.Returning{}).Where("id=? AND user_id=?", diaryId, userId).Updates(model.Diary{
		Title:       diary.Title,
		Description: diary.Description,
		Weather:     diary.Weather,
		Date:        diary.Date,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *diaryRepository) DeleteDiary(userId uint, diaryId uint) error {
	result := tr.db.Where("id=? AND user_id=?", diaryId, userId).Delete(&model.Diary{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
