package usecase_test

import (
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/mock"
)

func (suite *ArticleUsecaseTestSuite) TestIndex() {

	request := suite.initializeIndexRequest()
	responseMock := suite.initializeIndexResponse()

	sliceArticle := []models.Article{}

	suite.Run("will failed because data from cache is nil and postgres is failed", func() {

		suite.articleCache.On("Get", mock.Anything, mock.Anything).
			Return(nil, fiber.ErrNotFound).Once()

		suite.articleRepo.On("FindAll", mock.Anything, mock.Anything).
			Return(sliceArticle, fiber.ErrUnprocessableEntity).Once()

		data, err := suite.articleUsecase.Index(request)

		suite.Assert().NotNil(err)
		suite.Assert().Nil(data)
	})

	suite.Run("will failed because data cause fail get count", func() {

		suite.articleCache.On("Get", mock.Anything, mock.Anything).
			Return(nil, fiber.ErrNotFound).Once()

		suite.articleRepo.On("FindAll", mock.Anything, mock.Anything).
			Return(sliceArticle, nil).Once()

		suite.articleRepo.On("Count", mock.Anything, mock.Anything).
			Return(int64(0), fiber.ErrUnprocessableEntity).Once()

		data, err := suite.articleUsecase.Index(request)

		suite.Assert().NotNil(err)
		suite.Assert().Nil(data)
	})

	suite.Run("will success get  by cache", func() {

		suite.articleCache.On("Get", mock.Anything, mock.Anything).
			Return(responseMock, nil).Once()

		data, err := suite.articleUsecase.Index(request)

		suite.Assert().NotNil(data)
		suite.Assert().Nil(err)
	})

	suite.Run("will success cache is nil and database is not nil", func() {

		suite.articleCache.On("Get", mock.Anything, mock.Anything).
			Return(nil, fiber.ErrUnprocessableEntity).Once()

		suite.articleRepo.On("FindAll", mock.Anything, mock.Anything).
			Return(responseMock.Data, nil).Once()

		suite.articleRepo.On("Count", mock.Anything, mock.Anything).
			Return(responseMock.Meta.Total, fiber.ErrUnprocessableEntity).Once()

		suite.articleCache.On("Set", mock.Anything, mock.Anything).
			Return(nil).Once()

		suite.Assert().Equal(int64(len(responseMock.Data)), responseMock.Meta.Total)

		meta := utils.PaginationMeta{
			Offset: request.Offset(),
			Limit:  request.Limit(),
			Total:  0,
		}

		meta.Total = responseMock.Meta.Total

		// response := response.Index{
		// 	Data: responseMock.Data,
		// 	Meta: meta,
		// }

		// expectedResponse := &response

		data, err := suite.articleUsecase.Index(request)

		suite.Assert().NotNil(err)
		suite.Assert().Nil(data)

		// check response
		// suite.Assert().Equal(expectedResponse, data)

	})

}

func (suite *ArticleUsecaseTestSuite) initializeIndexRequest() utils.PaginationConfig {

	var request utils.PaginationConfig

	filterable := map[string]string{
		"title": utils.StringType,
		"body":  utils.StringType,
	}

	conditions := map[string][]string{
		"sort": {"created_at DESC"},
	}

	request = utils.NewRequestPaginationConfig(conditions, filterable)

	return request

}

func (suite *ArticleUsecaseTestSuite) initializeIndexResponse() *response.Index {

	data := []models.Article{
		{
			ID:        1,
			Author:    "Khalid",
			Title:     "Football",
			Body:      "football, also called association football or soccer, game in which two teams of 11 players, using any part of their bodies except their hands and arms, try to maneuver the ball into the opposing team’s goal. Only the goalkeeper is permitted to handle the ball and may do so only within the penalty area surrounding the goal. The team that scores more goals wins",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Author:    "Khalid",
			Title:     "IDS",
			Body:      "An Intrusion Detection System (IDS) is a monitoring system that detects suspicious activities and generates alerts when they are detected. Based upon these alerts, a security operations center (SOC) analyst or incident responder can investigate the issue and take the appropriate actions to remediate the threat",
			CreatedAt: time.Now(),
		},
	}

	meta := utils.PaginationMeta{
		Limit:  20,
		Offset: 0,
		Total:  2,
	}

	return &response.Index{
		Data: data,
		Meta: meta,
	}

}
