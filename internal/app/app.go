package app

import (
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	noteV1 "github.com/impopov/note-server-api/internal/app/api/note_v1"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// App ..
type App struct {
	serviceProvider *serviceProvider
	noteImpl        *noteV1.Implementation

	pathConfig string

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

// NewApp ..
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

// Run ..
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := a.runGRPC(wg)
	if err != nil {
		return err
	}

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().GRPC.GetAddressGRPC())
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to process gRPC server: %s", err.Error())
		}
	}()

	log.Printf("Run gRPC server on %s port\n", a.serviceProvider.GetConfig().GRPC.GetAddressGRPC())
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(a.serviceProvider.GetConfig().HTTP.GetAddressHTTP(), a.mux); err != nil {
			log.Fatalf("failed to process muxer: %s", err.Error())
		}
	}()

	log.Printf("Run public http handler on %s port\n", a.serviceProvider.GetConfig().HTTP.GetAddressHTTP())
	return nil
}

func (a *App) initDeps(ctx context.Context) error {

	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = noteV1.NewImplementation(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteServiceV1Server(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterNoteServiceV1HandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().GRPC.GetAddressGRPC(), opts)
	if err != nil {
		return err
	}

	return nil
}
