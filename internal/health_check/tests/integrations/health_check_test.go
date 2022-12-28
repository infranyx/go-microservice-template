package artcileIntegrationTest

import (
	"context"
	"encoding/json"
	healthCheckDto "github.com/infranyx/go-grpc-template/internal/health_check/dto"
	healthCheckFixture "github.com/infranyx/go-grpc-template/internal/health_check/tests/fixtures"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testSuite struct {
	suite.Suite
	fixture *healthCheckFixture.IntegrationTestFixture
}

func (suite *testSuite) SetupSuite() {
	fixture, err := healthCheckFixture.NewIntegrationTestFixture()
	if err != nil {
		assert.Error(suite.T(), err)
	}

	suite.fixture = fixture
}

func (suite *testSuite) TearDownSuite() {
	suite.fixture.TearDown()
}

func (suite *testSuite) TestHealthCheckHttpShouldSendOkForAllUnits() {

	request := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()
	suite.fixture.InfraContainer.EchoHttpServer.GetEchoInstance().ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)

	healthCheckResponseDto := new(healthCheckDto.HealthCheckResponseDto)

	if assert.NoError(suite.T(), json.Unmarshal(response.Body.Bytes(), healthCheckResponseDto)) {
		assert.Equal(suite.T(), true, healthCheckResponseDto.Status)
		assert.Equal(suite.T(), []healthCheckDto.HealthCheckUnit{
			{
				Unit: "postgres",
				Up:   true,
			},
			{
				Unit: "kafka",
				Up:   true,
			},
			{
				Unit: "writable-tmp-dir",
				Up:   true,
			},
		}, healthCheckResponseDto.Units)
	}

}

func (suite *testSuite) TestHealthCheckGrpcShouldSendServingForAllServices() {
	ctx := context.Background()

	healthCheckRequest := &grpcHealthV1.HealthCheckRequest{
		Service: "all",
	}
	response, _ := suite.fixture.HealthCheckGrpcClient.Check(ctx, healthCheckRequest)

	assert.NotNil(suite.T(), response)
	assert.Equal(suite.T(), grpcHealthV1.HealthCheckResponse_SERVING, response.GetStatus())
}

func (suite *testSuite) TestHealthCheckGrpcShouldSendUnknownForUnknownService() {
	ctx := context.Background()

	healthCheckRequest := &grpcHealthV1.HealthCheckRequest{
		Service: "un-known-service",
	}
	response, _ := suite.fixture.HealthCheckGrpcClient.Check(ctx, healthCheckRequest)

	assert.NotNil(suite.T(), response)
	assert.Equal(suite.T(), grpcHealthV1.HealthCheckResponse_UNKNOWN, response.GetStatus())
}

func TestRunSuite(t *testing.T) {
	tSuite := new(testSuite)
	suite.Run(t, tSuite)
}
