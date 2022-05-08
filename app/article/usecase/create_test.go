package usecase_test

import (
	"backend/app/article/delivery/http/request"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

func (suite *ArticleUsecaseTestSuite) TestCreate() {

	request := suite.initializeCreateRequest()

	suite.Run("fail on article insert", func() {

		suite.articleRepo.On("Insert", mock.Anything, mock.Anything).
			Return(fiber.ErrUnprocessableEntity).Once()

		data, err := suite.articleUsecase.Create(request)
		suite.Assert().NotNil(err)
		suite.Assert().Nil(data)

	})

	suite.Run("success run without error", func() {

		suite.articleRepo.On("Insert", mock.Anything, mock.Anything).
			Return(nil).Once()

		data, err := suite.articleUsecase.Create(request)
		suite.Assert().Nil(err)
		suite.Assert().NotNil(data)


		// using  go routine
		suite.articleCache.On("FlushAll", mock.Anything, mock.Anything).
		Return(nil).Once()

	})

}

func (suite *ArticleUsecaseTestSuite) initializeCreateRequest() request.Create {
	return request.Create{
		Author: "Khalid",
		Title:  "Football",
		Body:   "football, also called association football or soccer, game in which two teams of 11 players, using any part of their bodies except their hands and arms, try to maneuver the ball into the opposing team’s goal. Only the goalkeeper is permitted to handle the ball and may do so only within the penalty area surrounding the goal. The team that scores more goals wins.",
	}
}
