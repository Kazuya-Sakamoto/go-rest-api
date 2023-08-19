package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	GetAllArticles(articles *[]model.Article) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) GetAllArticles(articles *[]model.Article) error {
	if err := ar.db.Order("created_at").Find(articles).Error; err != nil {
		return err
	}
	return nil
}
