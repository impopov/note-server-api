package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	noteV1 "github.com/impopov/note-server-api/internal/app/api/note_v1"
	"github.com/impopov/note-server-api/internal/repository"
	"github.com/impopov/note-server-api/internal/service/note"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

const (
	hostGrpc = "localhost:50051"
	hostHttp = "localhost:8090"

	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		startGRPC()
	}()

	go func() {
		defer wg.Done()
		startHttp()
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
		return err
	}

	dbDSN := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := repository.NewRepository(db)
	noteService := note.NewService(noteRepository)

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	desc.RegisterNoteServiceV1Server(s, noteV1.NewImplementation(noteService))

	fmt.Println("Starting GRPC server on port", hostGrpc)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve:%s", err.Error())
		return err
	}

	return nil
}

func startHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteServiceV1HandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}

	fmt.Println("Starting http server on port", hostHttp)

	return http.ListenAndServe(hostHttp, mux)
}
