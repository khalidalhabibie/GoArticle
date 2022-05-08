package http_test

import (
	"backend/app/article/delivery/http/request"
	"backend/app/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

func (suite *ArticleHttpTestSuite) TestCreate() {
	// requestSucece := suite.initializeCreateRequest()
	suite.App.Post("/articles", suite.articleHandler.Create)

	suite.Run("will failed because parse request", func() {

		resp, err := suite.App.Test(httptest.NewRequest(fiber.MethodPost, "/articles", nil))

		// check test app test error
		suite.Assert().Equal(err, nil)
		suite.Assert().Equal(fiber.ErrBadRequest.Code, resp.StatusCode)

	})
	suite.Run("will failed because validate request", func() {

		request := request.Create{}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(request)
		if err != nil {
			log.Fatal(err)
		}

		req := httptest.NewRequest(fiber.MethodPost, "/articles", &buf)
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.App.Test(req)

		// check test app test error
		suite.Assert().Equal(err, nil)
		suite.Assert().Equal(fiber.ErrBadRequest.Code, resp.StatusCode)

	})
	suite.Run("will fail because usecase (bisnis logic or repository) failed", func() {

		request := suite.initializeCreateRequest()

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(request)
		if err != nil {
			log.Fatal(err)
		}

		suite.articleUsecase.On("Create", mock.Anything, mock.Anything).
			Return(nil, fiber.ErrUnprocessableEntity).Once()

		req := httptest.NewRequest(fiber.MethodPost, "/articles", &buf)
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.App.Test(req)

		// check test app test error
		suite.Assert().Equal(err, nil)
		suite.Assert().Equal(fiber.ErrUnprocessableEntity.Code, resp.StatusCode)

	})
	suite.Run("Sucess", func() {

		request := suite.initializeCreateRequest()

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(request)
		if err != nil {
			log.Fatal(err)
		}

		response := &models.Article{
			ID:        1,
			Author:    request.Author,
			Body:      request.Body,
			CreatedAt: time.Now(),
		}

		suite.articleUsecase.On("Create", mock.Anything, mock.Anything).
			Return(response, nil).Once()

		req := httptest.NewRequest(fiber.MethodPost, "/articles", &buf)
		req.Header.Set("Content-Type", "application/json")

		resp, err := suite.App.Test(req)

		// check test app test error
		suite.Assert().Equal(err, nil)
		suite.Assert().Equal(fiber.StatusOK, resp.StatusCode)

	})

}

func (suite *ArticleHttpTestSuite) initializeCreateRequest() request.Create {
	return request.Create{
		Author: "Khalid",
		Title:  "Football",
		Body:   "football, also called association football or soccer, game in which two teams of 11 players, using any part of their bodies except their hands and arms, try to maneuver the ball into the opposing team’s goal. Only the goalkeeper is permitted to handle the ball and may do so only within the penalty area surrounding the goal. The team that scores more goals wins.",
	}
}
