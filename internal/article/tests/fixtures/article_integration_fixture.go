package fixture

import (
	"context"
	"math"
	"net"
	"time"

	goTemplateUseCase "github.com/infranyx/go-grpc-template/external/go_template/usecase"
	articleGrpc "github.com/infranyx/go-grpc-template/internal/article/delivery/grpc"
	articleHttp "github.com/infranyx/go-grpc-template/internal/article/delivery/http"
	articleKafkaProducer "github.com/infranyx/go-grpc-template/internal/article/delivery/kafka/producer"
	articleRepo "github.com/infranyx/go-grpc-template/internal/article/repository"
	articleUseCase "github.com/infranyx/go-grpc-template/internal/article/usecase"
	clientContainer "github.com/infranyx/go-grpc-template/pkg/client_container"
	iContainer "github.com/infranyx/go-grpc-template/pkg/infra_container"
	"github.com/infranyx/go-grpc-template/pkg/logger"
	articleV1 "github.com/infranyx/protobuf-template-go/golang-grpc-template/article/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type IntegrationTestFixture struct {
	InfraContainer    *iContainer.IContainer
	Ctx               context.Context
	cancel            context.CancelFunc
	Cleanup           func()
	ArticleGrpcClient articleV1.ArticleServiceClient
}

func NewIntegrationTestFixture() (*IntegrationTestFixture, error) {
	deadline := time.Now().Add(time.Duration(math.MaxInt64))
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	ic, cleanup, err := iContainer.NewIC(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	client, clientCleanup, err := clientContainer.NewCC(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	goTempUseCase := goTemplateUseCase.NewGoTemplateUseCase(client.GoTemplateGrpc)
	repo := articleRepo.NewArticleRepository(ic.Pg)
	producer := articleKafkaProducer.NewArticleProducer(ic.KafkaWriter)
	usecase := articleUseCase.NewArticleUseCase(repo, goTempUseCase, producer)

	// echo
	ic.EchoServer.SetupDefaultMiddlewares()
	groupAPI := ic.EchoServer.GetEchoInstance().Group(ic.EchoServer.GetBasePath())
	echoController := articleHttp.NewArticleHttpController(usecase)
	articleHttp.NewArticleAPI(echoController).Register(groupAPI)

	// grpc
	grpcContrller := articleGrpc.NewArticleGrpcController(usecase)

	lis := bufconn.Listen(bufSize)
	articleV1.RegisterArticleServiceServer(ic.GrpcServer.GetCurrentGrpcServer(), grpcContrller)
	go func() {
		if err := ic.GrpcServer.GetCurrentGrpcServer().Serve(lis); err != nil {
			logger.Zap.Sugar().Fatalf("Server exited with error: %v", err)
		}
	}()

	// init article grpc client environment
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		cancel()
		return nil, err
	}
	articleGrpcClient := articleV1.NewArticleServiceClient(conn)

	return &IntegrationTestFixture{
		Cleanup: func() {
			cancel()
			cleanup()
			conn.Close()
			clientCleanup()
		},
		InfraContainer:    ic,
		Ctx:               ctx,
		cancel:            cancel,
		ArticleGrpcClient: articleGrpcClient,
	}, nil
}
