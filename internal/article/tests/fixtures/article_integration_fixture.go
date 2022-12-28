package articleFixture

import (
	"context"
	"math"
	"net"
	"time"

	sampleExtServiceUseCase "github.com/infranyx/go-grpc-template/external/sample_ext_service/usecase"
	articleGrpc "github.com/infranyx/go-grpc-template/internal/article/delivery/grpc"
	articleHttp "github.com/infranyx/go-grpc-template/internal/article/delivery/http"
	articleKafkaProducer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/producer"
	articleRepo "github.com/infranyx/go-grpc-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	externalBridge "github.com/infranyx/go-grpc-template/pkg/external_bridge"
	iContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const BUFSIZE = 1024 * 1024

type IntegrationTestFixture struct {
	TearDown          func()
	Ctx               context.Context
	Cancel            context.CancelFunc
	InfraContainer    *iContainer.IContainer
	ArticleGrpcClient articleV1.ArticleServiceClient
}

func NewIntegrationTestFixture() (*IntegrationTestFixture, error) {
	deadline := time.Now().Add(time.Duration(math.MaxInt64))
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	ic, infraDown, err := iContainer.NewIC(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	extBridge, extBridgeDown, err := externalBridge.NewExternalBridge(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(extBridge.SampleExtGrpcService)
	kafkaProducer := articleKafkaProducer.NewProducer(ic.KafkaWriter)
	repository := articleRepo.NewRepository(ic.Postgres)
	useCase := articleUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// http
	ic.EchoHttpServer.SetupDefaultMiddlewares()
	httpRouterGp := ic.EchoHttpServer.GetEchoInstance().Group(ic.EchoHttpServer.GetBasePath())
	httpController := articleHttp.NewController(useCase)
	articleHttp.NewRouter(httpController).Register(httpRouterGp)

	// grpc
	grpcController := articleGrpc.NewController(useCase)
	articleV1.RegisterArticleServiceServer(ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	lis := bufconn.Listen(BUFSIZE)
	go func() {
		if err := ic.GrpcServer.GetCurrentGrpcServer().Serve(lis); err != nil {
			logger.Zap.Sugar().Fatalf("Server exited with error: %v", err)
		}
	}()

	grpcClientConn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	articleGrpcClient := articleV1.NewArticleServiceClient(grpcClientConn)

	return &IntegrationTestFixture{
		TearDown: func() {
			cancel()
			infraDown()
			_ = grpcClientConn.Close()
			extBridgeDown()
		},
		InfraContainer:    ic,
		Ctx:               ctx,
		Cancel:            cancel,
		ArticleGrpcClient: articleGrpcClient,
	}, nil
}
