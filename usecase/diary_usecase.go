package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IDiaryUsecase interface {
	GetAllDiaries(userId uint) ([]model.DiaryResponse, error)
}

type diaryUsecase struct {
	dr repository.IDiaryRepository
}

func NewDiaryUsecase(dr repository.IDiaryRepository) IDiaryUsecase {
	return &diaryUsecase{dr}
}

func (du *diaryUsecase) GetAllDiaries(userId uint) ([]model.DiaryResponse, error) {
	diaries := []model.Diary{}
	if err := du.dr.GetAllDiaries(&diaries, userId); err != nil {
		return nil, err
	}
	responseDiaries := []model.DiaryResponse{}
	for _, v := range diaries {
		d := model.DiaryResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Weather:     v.Weather,
			Date:        v.Date,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		responseDiaries = append(responseDiaries, d)
	}
	return responseDiaries, nil
}
