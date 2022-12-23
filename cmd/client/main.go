package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceV1Client(con)
	ctx := context.Background()

	// Menu-based options allowing the user to choose from CRUD
	fmt.Println("Enter the one of the following choices below:")
	fmt.Print("1 - Create Note\n2 - Get note by id\n3 - Get All notes\n4 - Update Note gy id\n5 - Delete Note by id\n")

	choice := bufio.NewReader(os.Stdin)
	text, err := choice.ReadString('\n')
	if err != nil {
		log.Fatalf("Can' read input: %s", err.Error())
	}

	switch text {
	case "1\n":
		//Client call CreateNote
		res, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
			Title:  "Tom Sawyer and Huckleberry Finn",
			Text:   "Right is right, and wrong is wrong, and a body ain't got no business doing wrong when he ain't ignorant and knows better.",
			Author: "Mark Twain",
		})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("Create Note")
		log.Println("Id:", res.GetId())

	case "2\n":
		//Client call GetNote
		res, err := client.GetNote(ctx, &desc.GetNoteRequest{Id: 1})
		if err != nil {
			log.Println(err.Error())
		}

		log.Printf("Note by id %d is:", res.GetNote().GetId())
		fmt.Println("Title:", res.GetNote().GetTitle())
		fmt.Println("Text:", res.GetNote().GetText())
		fmt.Println("Author:", res.GetNote().GetAuthor())
		fmt.Println("Updated at:", res.GetNote().GetUpdatedAt().AsTime())

	case "3\n":
		//Client call GetAllNote
		res, err := client.GetListNote(ctx, &desc.Empty{})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("All Notes:")
		fmt.Println(res.GetNote())

	case "4\n":
		//Client call UpdateNote
		_, err := client.UpdateNote(ctx, &desc.UpdateNoteRequest{Note: &desc.Note{
			Id:     6,
			Title:  "New Title",
			Text:   "New Text",
			Author: "New Author",
		}})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("Update Note")

	case "5\n":
		//Client call DeleteNote
		_, err := client.DeleteNote(ctx, &desc.DeleteNoteRequest{Id: 4})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("Delete Note")

	default:
		fmt.Println("\nWrong option!")
	}
}
