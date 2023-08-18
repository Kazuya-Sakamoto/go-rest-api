package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IDiaryCommentRepository interface {
	GetDiaryCommentsByDiaryIDAndUserID(diaryComments *[]model.DiaryComment, diaryId, userId uint) error
	CreateDiaryComment(diaryComment *model.DiaryComment) error
	DeleteDiaryComment(userId, diaryCommentId uint) error
}

type diaryCommentRepository struct {
	db *gorm.DB
}

func NewDiaryCommentRepository(db *gorm.DB) IDiaryCommentRepository {
	return &diaryCommentRepository{db}
}

func (dcr *diaryCommentRepository) GetDiaryCommentsByDiaryIDAndUserID(diaryComments *[]model.DiaryComment, diaryId, userId uint) error {
	if err := dcr.db.
		Preload("Diary").
		Preload("User").
		Where("diary_comments.diary_id = ? AND diary_comments.user_id = ?", diaryId, userId).
		Order("diary_comments.created_at").
		Find(diaryComments).Error; err != nil {
		return err
	}
	return nil
}

func (dcr *diaryCommentRepository) CreateDiaryComment(diaryComment *model.DiaryComment) error {
	if err := dcr.db.Create(diaryComment).Error; err != nil {
		return err
	}
	return nil
}

func (dcr *diaryCommentRepository) DeleteDiaryComment(userId, diaryCommentId uint) error {
	result := dcr.db.Where("id=? AND user_id=?", diaryCommentId, userId).Delete(&model.DiaryComment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
