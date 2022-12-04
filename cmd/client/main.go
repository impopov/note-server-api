package main

import (
	"bufio"
	"context"
	"fmt"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
	"os"
)

const address = "localhost:50051"

func main() {
	//Create client connection
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)
	ctx := context.Background()

	// Menu-based options allowing the user to choose from CRUD
	fmt.Println("Enter the one of the following choices below:")
	fmt.Print("1 - Create Note\n2 - Get note by id\n3 - Get All notes\n4 - Update Note gy id\n5 - Delete Note by id\n")

	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	switch text {
	case "1\n":
		//Client call CreateNote
		resCreateNote, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
			Title:  "Tom Sawyer and Huckleberry Finn",
			Text:   "Right is right, and wrong is wrong, and a body ain't got no business doing wrong when he ain't ignorant and knows better.",
			Author: "Mark Twain",
		})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("Create Note")
		log.Println("Id:", resCreateNote.Id)

	case "2\n":
		//Client call GetNote
		resGetNote, err := client.GetNote(ctx, &desc.GetNoteRequest{Id: 2})
		if err != nil {
			log.Println(err.Error())
		}

		log.Printf("Note by id %d is:", resGetNote.GetNote().GetId())
		fmt.Println("Title:", resGetNote.Note.GetTitle())
		fmt.Println("Text:", resGetNote.Note.GetText())
		fmt.Println("Author:", resGetNote.Note.GetAuthor())

	case "3\n":
		//Client call GetAllNote
		resGetAllNote, err := client.GetAllNote(ctx, &desc.GetAllNoteRequest{})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("All Notes:")
		fmt.Println(resGetAllNote)

	case "4\n":
		//Client call UpdateNote
		resUpdateNote, err := client.UpdateNote(ctx, &desc.UpdateNoteRequest{Note: &desc.Note{
			Id:     1,
			Title:  "New Title",
			Text:   "New Text",
			Author: "New Author",
		}})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("Update Note")
		log.Println("Status:", resUpdateNote.GetStatus())

	case "5\n":
		//Client call DeleteNote
		resDeleteNote, err := client.DeleteNote(ctx, &desc.DeleteNoteRequest{Id: 1})
		if err != nil {
			log.Println(err.Error())
		}

		log.Println("Delete Note")
		log.Println("Status:", resDeleteNote.GetStatus())

	default:
		fmt.Println("\nWrong option!")
	}
}
