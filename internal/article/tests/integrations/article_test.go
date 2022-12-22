package artcileIntegrationTest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	fixture "github.com/infranyx/go-grpc-template/internal/article/tests/fixtures"
	grpcError "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"
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
	articleJSON := `{"name":"John Snow","desc":"King of the north"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/article", strings.NewReader(articleJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	suite.fixture.InfraContainer.EchoServer.GetEchoInstance().ServeHTTP(rec, req)

	// Assertions
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	articleDto := new(articleDto.Article)
	if assert.NoError(suite.T(), json.Unmarshal(rec.Body.Bytes(), articleDto)) {
		assert.NotNil(suite.T(), articleDto.ID)
		assert.Equal(suite.T(), "John Snow", articleDto.Name)
		assert.Equal(suite.T(), "King of the north", articleDto.Description)
	}

}

func (suite *ArticleSuiteTests) TestNameValidationErrCreateHttpArticle() {
	articleJSON := `{"name":"Jo","desc":"King of the north"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/article", strings.NewReader(articleJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	suite.fixture.InfraContainer.EchoServer.SetupDefaultMiddlewares()
	suite.fixture.InfraContainer.EchoServer.GetEchoInstance().ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	httpErr := httpError.ParseExternalHttpErr(rec.Result().Body)
	if assert.NotNil(suite.T(), httpErr) {
		assert.Equal(suite.T(), http.StatusBadRequest, httpErr.GetStatus())
		assert.Contains(suite.T(), httpErr.GetDetails(), "name")
	}

}

func (suite *ArticleSuiteTests) TestDescValidationErrCreateHttpArticle() {
	articleJSON := `{"name":"John Snow","desc":"King"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/article", strings.NewReader(articleJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	suite.fixture.InfraContainer.EchoServer.SetupDefaultMiddlewares()
	suite.fixture.InfraContainer.EchoServer.GetEchoInstance().ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	httpErr := httpError.ParseExternalHttpErr(rec.Result().Body)
	if assert.NotNil(suite.T(), httpErr) {
		assert.Equal(suite.T(), http.StatusBadRequest, httpErr.GetStatus())
		assert.Contains(suite.T(), httpErr.GetDetails(), "desc")
	}
}

func TestRunSuite(t *testing.T) {
	suiteTester := new(ArticleSuiteTests)
	suite.Run(t, suiteTester)
}
