package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IDiaryCommentUsecase interface {
	GetDiaryCommentsByDiaryIDAndUserID(diaryId, userId uint) ([]model.DiaryCommentResponse, error)
	CreateDiaryComment(diaryComment model.DiaryComment) (model.DiaryCommentResponse, error)
	DeleteDiaryComment(userId, diaryCommentId uint) error
}

type diaryCommentUsecase struct {
	dcr repository.IDiaryCommentRepository
}

func NewDiaryCommentRepository(dcr repository.IDiaryCommentRepository) IDiaryCommentUsecase {
	return &diaryCommentUsecase{dcr}
}

func (dcu *diaryCommentUsecase) GetDiaryCommentsByDiaryIDAndUserID(diaryId, userId uint) ([]model.DiaryCommentResponse, error) {
	diaryComments := []model.DiaryComment{}
	if err := dcu.dcr.GetDiaryCommentsByDiaryIDAndUserID(&diaryComments, diaryId, userId); err != nil {
		return nil, err
	}
	responseDiaryCommnets := []model.DiaryCommentResponse{}
	for _, v := range diaryComments {
		dc := model.DiaryCommentResponse{
			ID:        v.ID,
			Comment:   v.Comment,
			Diary:     v.Diary,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		responseDiaryCommnets = append(responseDiaryCommnets, dc)
	}
	return responseDiaryCommnets, nil
}

func (dcu *diaryCommentUsecase) CreateDiaryComment(diaryComment model.DiaryComment) (model.DiaryCommentResponse, error) {
	if err := dcu.dcr.CreateDiaryComment(&diaryComment); err != nil {
		return model.DiaryCommentResponse{}, err
	}
	responseDiaryCommnet := model.DiaryCommentResponse{
		ID:        diaryComment.ID,
		Comment:   diaryComment.Comment,
		Diary:     diaryComment.Diary,
		CreatedAt: diaryComment.CreatedAt,
		UpdatedAt: diaryComment.UpdatedAt,
	}
	return responseDiaryCommnet, nil
}

func (dcu *diaryCommentUsecase) DeleteDiaryComment(userId, diaryCommentId uint) error {
	if err := dcu.dcr.DeleteDiaryComment(userId, diaryCommentId); err != nil {
		return err
	}
	return nil
}
