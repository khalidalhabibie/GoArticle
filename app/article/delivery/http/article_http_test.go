package http_test

import (
	articlehHttp "backend/app/article/delivery/http"
	"backend/app/article/mocks"
	"backend/pkg/configs"
	"testing"

	mocket "github.com/Selvatico/go-mocket"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type ArticleHttpTestSuite struct {
	suite.Suite

	articleUsecase *mocks.Usecase
	articleHandler articlehHttp.Handler

	App *fiber.App
}

func TestArticleHttp(t *testing.T) {

	suite.Run(t, new(ArticleHttpTestSuite))

}

func (suite *ArticleHttpTestSuite) SetupTest() {

	suite.articleUsecase = new(mocks.Usecase)

	suite.articleHandler = *articlehHttp.New(suite.articleUsecase)

	mocket.Catcher.Register()

	mocket.Catcher.Logging = true

	// database.SetupDBTests()

	// Define Fiber config
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	suite.App = fiber.New(config)

}
