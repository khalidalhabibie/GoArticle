package models

const (

	// layer
	LogLayerDelivery   = "delivery"
	LogLayerUsecase    = "usecase"
	LogLayerRepository = "repository"

	// service
	ServiceAuth    = "auth"
	ServiceArticle = "article"

	// error type
	LogErrorTypeConnectionDatabase = "database connection"
	LogErrorTypeConnectionRedis    = "redis connection"
)
