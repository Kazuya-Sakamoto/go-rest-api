package usecase

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IDiaryUsecase interface {
	GetAllDiaries(userId uint) ([]model.DiaryResponse, error)
	GetDiaryById(userId uint, diaryId uint) (model.DiaryResponse, error)
	CreateDiary(diary model.Diary) (model.DiaryResponse, error)
	UpdateDiary(diary model.Diary, userId uint, diaryId uint) (model.DiaryResponse, error)
	DeleteDiary(userId uint, diaryId uint) error
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

func (du *diaryUsecase) GetDiaryById(userId uint, diaryId uint) (model.DiaryResponse, error) {
	diary := model.Diary{}
	if err := du.dr.GetDiaryById(&diary, userId, diaryId); err != nil {
		return model.DiaryResponse{}, err
	}
	responseDiary := model.DiaryResponse{
		ID:          diary.ID,
		Title:       diary.Title,
		Description: diary.Description,
		Weather:     diary.Weather,
		Date:        diary.Date,
		CreatedAt:   diary.CreatedAt,
		UpdatedAt:   diary.UpdatedAt,
	}
	return responseDiary, nil
}

func (du *diaryUsecase) CreateDiary(diary model.Diary) (model.DiaryResponse, error) {
	if err := du.dr.CreateDiary(&diary); err != nil {
		return model.DiaryResponse{}, err
	}

	responseDiary := model.DiaryResponse{
		ID:          diary.ID,
		Title:       diary.Title,
		Description: diary.Description,
		Weather:     diary.Weather,
		Date:        diary.Date,
		CreatedAt:   diary.CreatedAt,
		UpdatedAt:   diary.UpdatedAt,
	}
	fmt.Println(responseDiary, "usecase/diary_usecase.go > responseDiary")
	return responseDiary, nil
}

func (du *diaryUsecase) UpdateDiary(diary model.Diary, userId uint, diaryId uint) (model.DiaryResponse, error) {
	if err := du.dr.UpdateDiary(&diary, userId, diaryId); err != nil {
		return model.DiaryResponse{}, err
	}
	responseDiary := model.DiaryResponse{
		ID:          diary.ID,
		Title:       diary.Title,
		Description: diary.Description,
		Weather:     diary.Weather,
		Date:        diary.Date,
		CreatedAt:   diary.CreatedAt,
		UpdatedAt:   diary.UpdatedAt,
	}
	return responseDiary, nil
}

func (du *diaryUsecase) DeleteDiary(userId uint, diaryId uint) error {
	if err := du.dr.DeleteDiary(userId, diaryId); err != nil {
		return err
	}
	return nil
}
