package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IArticleUsecase interface {
	GetAllArticles() ([]model.ArticleResponse, error)
}

type articleUsecase struct {
	ar repository.IArticleRepository
}

func NewArticleUsecase(ar repository.IArticleRepository) IArticleUsecase {
	return &articleUsecase{ar}
}

func (au *articleUsecase) GetAllArticles() ([]model.ArticleResponse, error) {
	articles := []model.Article{}
	if err := au.ar.GetAllArticles(&articles); err != nil {
		return nil, err
	}
	response := []model.ArticleResponse{}
	for _, v := range articles {
		a := model.ArticleResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		response = append(response, a)
	}
	return response, nil
}
