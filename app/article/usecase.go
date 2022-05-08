package article

import (
	"backend/app/article/delivery/http/request"
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"
)

type Usecase interface {
	Create(requsest request.Create) (*models.Article, error)
	// Index(request utils.PaginationConfig) ([]models.Article, utils.PaginationMeta, error)
	Index(request utils.PaginationConfig) (*response.Index, error)
	// Index(request utils.PaginationConfig) (interface{}, error)
}
