package main

import (
	"fmt"
	"log"
	"net"

	noteV1 "github.com/impopov/note-server-api/internal/app/api/note_v1"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteServiceV1Server(s, noteV1.NewImplementation())

	fmt.Println("Starting server on port", port)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve:%s", err.Error())
	}

}
