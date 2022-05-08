package article

import (
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(article *models.Article, tx *gorm.DB) error
	FindAll(config utils.PaginationConfig) ([]models.Article, error)
	Count(config utils.PaginationConfig) (int64, error)
}

type Cache interface {
	Set(request utils.PaginationConfig, response response.Index) error
	Get(request utils.PaginationConfig) (*response.Index, error)
	FlushAll() error
}
