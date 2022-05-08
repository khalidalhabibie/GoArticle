package http_test

import (
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"
	"net/http/httptest"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

func (suite *ArticleHttpTestSuite) TestIndex() {
	response := suite.initializeIndexRespone()
	suite.App.Get("/articles", suite.articleHandler.Index)

	suite.Run("will sucess with no param", func() {

		suite.articleUsecase.On("Index", mock.Anything, mock.Anything).
			Return(response, nil).Once()

		req := httptest.NewRequest(fiber.MethodGet, "/articles", nil)
		resp, err := suite.App.Test(req)

		// check error
		suite.Assert().Nil(err)
		suite.Assert().Equal(resp.StatusCode, fiber.StatusOK)

	})
	suite.Run("will sucess with param", func() {

		suite.articleUsecase.On("Index", mock.Anything, mock.Anything).
			Return(response, nil).Once()

		req := httptest.NewRequest(fiber.MethodGet, "/articles", nil)
		q := req.URL.Query()
		q.Add("limit", "1")
		q.Add("search", "chelsea")
		req.URL.RawQuery = q.Encode()

		resp, err := suite.App.Test(req)

		// check error
		suite.Assert().Nil(err)
		suite.Assert().Equal(resp.StatusCode, fiber.StatusOK)

	})
	suite.Run("will fail beacuse usecase fail", func() {

		suite.articleUsecase.On("Index", mock.Anything, mock.Anything).
			Return(nil, fiber.ErrUnprocessableEntity).Once()

		req := httptest.NewRequest(fiber.MethodGet, "/articles", nil)
		q := req.URL.Query()
		q.Add("limit", "1")
		q.Add("search", "chelsea")
		req.URL.RawQuery = q.Encode()

		resp, err := suite.App.Test(req)

		// check error
		suite.Assert().Nil(err)
		suite.Assert().Equal(resp.StatusCode, fiber.ErrUnprocessableEntity.Code)

	})

}

func (suite *ArticleHttpTestSuite) initializeIndexRespone() *response.Index {
	return &response.Index{
		Data: []models.Article{
			{
				ID:        1,
				Author:    "Khalid",
				Title:     "Football",
				Body:      "football, also called association football or soccer, game in which two teams of 11 players, using any part of their bodies except their hands and arms, try to maneuver the ball into the opposing team’s goal. Only the goalkeeper is permitted to handle the ball and may do so only within the penalty area surrounding the goal. The team that scores more goals wins.",
				CreatedAt: time.Now(),
			},
		},
		Meta: utils.PaginationMeta{
			Limit:  20,
			Offset: 0,
			Total:  1,
		},
	}
}
