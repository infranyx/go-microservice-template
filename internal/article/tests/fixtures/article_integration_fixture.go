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

	ic, iCCleanup, err := iContainer.NewIC(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	extBridge, extBridgeCleanUp, err := externalBridge.NewExternalBridge(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(extBridge.SampleExtGrpcService)
	kp := articleKafkaProducer.NewProducer(ic.KafkaWriter)
	rp := articleRepo.NewRepository(ic.Pg)
	uc := articleUseCase.NewUseCase(rp, seServiceUseCase, kp)

	// http
	ic.EchoServer.SetupDefaultMiddlewares()
	httpRouterGp := ic.EchoServer.GetEchoInstance().Group(ic.EchoServer.GetBasePath())
	httpController := articleHttp.NewController(uc)
	articleHttp.NewRouter(httpController).Register(httpRouterGp)

	// grpc
	grpcController := articleGrpc.NewController(uc)
	articleV1.RegisterArticleServiceServer(ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	lis := bufconn.Listen(BUFSIZE)
	go func() {
		if err := ic.GrpcServer.GetCurrentGrpcServer().Serve(lis); err != nil {
			logger.Zap.Sugar().Fatalf("Server exited with error: %v", err)
		}
	}()

	// init article grpc client environment
	grpcClientConn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		cancel()
		return nil, err
	}
	articleGrpcClient := articleV1.NewArticleServiceClient(grpcClientConn)

	return &IntegrationTestFixture{
		TearDown: func() {
			cancel()
			iCCleanup()
			grpcClientConn.Close()
			extBridgeCleanUp()
		},
		InfraContainer:    ic,
		Ctx:               ctx,
		Cancel:            cancel,
		ArticleGrpcClient: articleGrpcClient,
	}, nil
}
