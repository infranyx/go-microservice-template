package artcileIntegrationTest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
	"github.com/labstack/echo/v4"

	articleDto "github.com/infranyx/go-grpc-template/internal/article/dto"
	articleFixture "github.com/infranyx/go-grpc-template/internal/article/tests/fixtures"
	grpcError "github.com/infranyx/go-grpc-template/pkg/error/grpc"
	httpError "github.com/infranyx/go-grpc-template/pkg/error/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
)

type testSuite struct {
	suite.Suite
	fixture *articleFixture.IntegrationTestFixture
}

func (suite *testSuite) SetupSuite() {
	fixture, err := articleFixture.NewIntegrationTestFixture()
	if err != nil {
		assert.Error(suite.T(), err)
	}

	suite.fixture = fixture
}

func (suite *testSuite) TearDownSuite() {
	suite.fixture.TearDown()
}

func (suite *testSuite) TestSuccessfulCreateGrpcArticle() {
	ctx := context.Background()

	createArticleRequest := &articleV1.CreateArticleRequest{
		Name: "John",
		Desc: "Pro Developer",
	}

	response, err := suite.fixture.ArticleGrpcClient.CreateArticle(ctx, createArticleRequest)
	if err != nil {
		assert.Error(suite.T(), err)
	}

	assert.NotNil(suite.T(), response.Id)
	assert.Equal(suite.T(), "John", response.Name)
	assert.Equal(suite.T(), "Pro Developer", response.Desc)
}

func (suite *testSuite) TestNameValidationErrCreateGrpcArticle() {
	ctx := context.Background()

	createArticleRequest := &articleV1.CreateArticleRequest{
		Name: "Jo",
		Desc: "Pro Developer",
	}
	_, err := suite.fixture.ArticleGrpcClient.CreateArticle(ctx, createArticleRequest)

	assert.NotNil(suite.T(), err)

	grpcErr := grpcError.ParseExternalGrpcErr(err)
	assert.NotNil(suite.T(), grpcErr)
	assert.Equal(suite.T(), codes.InvalidArgument, grpcErr.GetStatus())
	assert.Contains(suite.T(), grpcErr.GetDetails(), "name")
}

func (suite *testSuite) TestDescValidationErrCreateGrpcArticle() {
	ctx := context.Background()

	createArticleRequest := &articleV1.CreateArticleRequest{
		Name: "John",
		Desc: "Pro",
	}
	_, err := suite.fixture.ArticleGrpcClient.CreateArticle(ctx, createArticleRequest)

	assert.NotNil(suite.T(), err)

	grpcErr := grpcError.ParseExternalGrpcErr(err)
	assert.NotNil(suite.T(), grpcErr)
	assert.Equal(suite.T(), codes.InvalidArgument, grpcErr.GetStatus())
	assert.Contains(suite.T(), grpcErr.GetDetails(), "desc")
}

func (suite *testSuite) TestSuccessCreateHttpArticle() {
	articleJSON := `{"name":"John Snow","desc":"King of the north"}`

	request := httptest.NewRequest(http.MethodPost, "/api/v1/article", strings.NewReader(articleJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()
	suite.fixture.InfraContainer.EchoHttpServer.GetEchoInstance().ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)

	caDto := new(articleDto.CreateArticleRequestDto)
	if assert.NoError(suite.T(), json.Unmarshal(response.Body.Bytes(), caDto)) {
		assert.Equal(suite.T(), "John Snow", caDto.Name)
		assert.Equal(suite.T(), "King of the north", caDto.Description)
	}

}

func (suite *testSuite) TestNameValidationErrCreateHttpArticle() {
	articleJSON := `{"name":"Jo","desc":"King of the north"}`

	request := httptest.NewRequest(http.MethodPost, "/api/v1/article", strings.NewReader(articleJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()
	suite.fixture.InfraContainer.EchoHttpServer.SetupDefaultMiddlewares()
	suite.fixture.InfraContainer.EchoHttpServer.GetEchoInstance().ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)

	httpErr := httpError.ParseExternalHttpErr(response.Result().Body)
	if assert.NotNil(suite.T(), httpErr) {
		assert.Equal(suite.T(), http.StatusBadRequest, httpErr.GetStatus())
		assert.Contains(suite.T(), httpErr.GetDetails(), "name")
	}

}

func (suite *testSuite) TestDescValidationErrCreateHttpArticle() {
	articleJSON := `{"name":"John Snow","desc":"King"}`

	request := httptest.NewRequest(http.MethodPost, "/api/v1/article", strings.NewReader(articleJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()

	suite.fixture.InfraContainer.EchoHttpServer.SetupDefaultMiddlewares()
	suite.fixture.InfraContainer.EchoHttpServer.GetEchoInstance().ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)

	httpErr := httpError.ParseExternalHttpErr(response.Result().Body)
	if assert.NotNil(suite.T(), httpErr) {
		assert.Equal(suite.T(), http.StatusBadRequest, httpErr.GetStatus())
		assert.Contains(suite.T(), httpErr.GetDetails(), "desc")
	}
}

func TestRunSuite(t *testing.T) {
	tSuite := new(testSuite)
	suite.Run(t, tSuite)
}
