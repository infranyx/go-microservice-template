package artcileIntegrationTest

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	fixture "github.com/infranyx/go-grpc-template/internal/article/tests/fixtures"
	"github.com/infranyx/go-grpc-template/pkg/config"
	grpcError "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"
	httpClient "github.com/infranyx/go-grpc-template/pkg/http/client"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
)

type ArticleSuiteTests struct {
	suite.Suite
	fixture *fixture.IntegrationTestFixture
}

func (suite *ArticleSuiteTests) SetupSuite() {
	fixture, err := fixture.NewIntegrationTestFixture()
	if err != nil {
		assert.Error(suite.T(), err)
	}
	suite.fixture = fixture
	if err != nil {
		assert.Error(suite.T(), err)
	}
}

func (suite *ArticleSuiteTests) TearDownSuite() {
	suite.fixture.Cleanup()
}

func (suite *ArticleSuiteTests) TestSuccessCreateGrpcArticle() {
	ctx := context.Background()
	input := &articleV1.CreateArticleRequest{
		Name: "John",
		Desc: "Pro Developer",
	}
	res, err := suite.fixture.ArticleGrpcClient.CreateArticle(ctx, input)
	if err != nil {
		assert.Error(suite.T(), err)
	}

	assert.NotNil(suite.T(), res.Id)
	assert.Equal(suite.T(), "John", res.Name)
	assert.Equal(suite.T(), "Pro Developer", res.Desc)
}

func (suite *ArticleSuiteTests) TestNameValidationErrCreateGrpcArticle() {
	ctx := context.Background()
	input := &articleV1.CreateArticleRequest{
		Name: "Jo",
		Desc: "Pro Developer",
	}
	_, err := suite.fixture.ArticleGrpcClient.CreateArticle(ctx, input)

	assert.NotNil(suite.T(), err)
	grpcErr := grpcError.ParseExternalGrpcErr(err)
	assert.NotNil(suite.T(), grpcErr)
	assert.Equal(suite.T(), codes.InvalidArgument, grpcErr.GetStatus())
	assert.Contains(suite.T(), grpcErr.GetDetails(), "name")
}

func (suite *ArticleSuiteTests) TestDescValidationErrCreateGrpcArticle() {
	ctx := context.Background()
	input := &articleV1.CreateArticleRequest{
		Name: "John",
		Desc: "Pro",
	}
	_, err := suite.fixture.ArticleGrpcClient.CreateArticle(ctx, input)

	assert.NotNil(suite.T(), err)

	grpcErr := grpcError.ParseExternalGrpcErr(err)
	assert.NotNil(suite.T(), grpcErr)
	assert.Equal(suite.T(), codes.InvalidArgument, grpcErr.GetStatus())
	assert.Contains(suite.T(), grpcErr.GetDetails(), "desc")
}

func (suite *ArticleSuiteTests) TestSuccessCreateHttpArticle() {
	ctx := context.Background()
	articleJSON := `{"name":"John Snow","desc":"King of the north"}`

	res, err := httpClient.
		BuildReq().
		SetContext(ctx).
		SetBody(strings.NewReader(articleJSON)).
		SetHeader(echo.HeaderContentType, echo.MIMEApplicationJSON).
		Post(fmt.Sprintf("http://localhost:%d%s/article", config.Conf.Http.Port, "/api/v1")).
		Execute()

	if err != nil {
		assert.Error(suite.T(), err)
	}

	articleDto := new(articleDto.Article)
	err = res.Bind(articleDto)
	if err != nil {
		assert.Error(suite.T(), err)
	}

	assert.NotNil(suite.T(), articleDto.ID)
	assert.Equal(suite.T(), "John Snow", articleDto.Name)
	assert.Equal(suite.T(), "King of the north", articleDto.Description)
}

func (suite *ArticleSuiteTests) TestNameValidationErrCreateHttpArticle() {
	ctx := context.Background()
	articleJSON := `{"name":"Jo","desc":"King of the north"}`

	res, err := httpClient.
		BuildReq().
		SetContext(ctx).
		SetBody(strings.NewReader(articleJSON)).
		SetHeader(echo.HeaderContentType, echo.MIMEApplicationJSON).
		Post(fmt.Sprintf("http://localhost:%d%s/article", config.Conf.Http.Port, "/api/v1")).
		Execute()

	if err != nil {
		assert.Error(suite.T(), err)
	}

	httpErr := httpError.ParseExternalHttpErr(res.Body())

	assert.NotNil(suite.T(), httpErr)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.GetStatus())
	assert.Equal(suite.T(), false, res.IsSuccess())
	assert.Contains(suite.T(), httpErr.GetDetails(), "name")
}

func (suite *ArticleSuiteTests) TestDescValidationErrCreateHttpArticle() {
	ctx := context.Background()
	articleJSON := `{"name":"John Snow","desc":"King"}`

	res, err := httpClient.
		BuildReq().
		SetContext(ctx).
		SetBody(strings.NewReader(articleJSON)).
		SetHeader(echo.HeaderContentType, echo.MIMEApplicationJSON).
		Post(fmt.Sprintf("http://localhost:%d%s/article", config.Conf.Http.Port, "/api/v1")).
		Execute()

	if err != nil {
		assert.Error(suite.T(), err)
	}

	httpErr := httpError.ParseExternalHttpErr(res.Body())

	assert.NotNil(suite.T(), httpErr)
	assert.Equal(suite.T(), http.StatusBadRequest, httpErr.GetStatus())
	assert.Equal(suite.T(), false, res.IsSuccess())
	assert.Contains(suite.T(), httpErr.GetDetails(), "desc")
}

func TestRunSuite(t *testing.T) {
	suiteTester := new(ArticleSuiteTests)
	suite.Run(t, suiteTester)
}
