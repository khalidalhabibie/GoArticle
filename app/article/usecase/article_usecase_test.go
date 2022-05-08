package usecase_test

import (
	"backend/app/article"
	"backend/app/article/mocks"
	"backend/app/article/usecase"
	"backend/platform/cache"
	"backend/platform/database"
	"testing"

	mocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
)

type ArticleUsecaseTestSuite struct {
	suite.Suite

	articleRepo  *mocks.Repository
	articleCache *mocks.Cache

	articleUsecase article.Usecase
}

func TestArticleUsecase(t *testing.T) {

	suite.Run(t, new(ArticleUsecaseTestSuite))

}

func (suite *ArticleUsecaseTestSuite) SetupTest() {
	suite.articleRepo = new(mocks.Repository)
	suite.articleCache = new(mocks.Cache)

	suite.articleUsecase = usecase.New(suite.articleRepo, suite.articleCache)

	mocket.Catcher.Register()

	mocket.Catcher.Logging = true

	database.SetupDBTests()
	cache.SetUpRedisForTesting()

	// suite.db = db

}
