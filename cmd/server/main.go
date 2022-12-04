package main

import (
	"fmt"
	"github.com/impopov/note-server-api/internal/app/api/note_v1"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	fmt.Println("Starting server on port", port)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve:%s", err.Error())
	}

}
