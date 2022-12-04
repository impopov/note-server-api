package main

import (
	"context"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)
	ctx := context.Background()

	//noteToWork := &desc.Note{
	//	Id:     1,
	//	Title:  "1",
	//	Text:   "2",
	//	Author: "3",
	//}

	res, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
		Title:  "Tom Sawyer and Huckleberry Finn",
		Text:   "“Right is right, and wrong is wrong, and a body ain't got no business doing wrong when he ain't ignorant and knows better.",
		Author: "Mark Twain",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", res.Id)
}
