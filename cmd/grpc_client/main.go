package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	desc "github.com/MAPiryazev/Auth/pkg/note_v1"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := desc.NewNoteV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getRequest, err := client.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Println("failed to get node by id", err)
	}
	println(getRequest.Name)

	createRequest, err := client.Create(ctx, &desc.CreateRequest{Info: &desc.NoteINfo{Name: "Ivan",
		Email:            "example@gmail.com",
		Password:         "1234",
		PasswordCoonfirm: "1234",
	}})
	if err != nil {
		log.Println("failed to create node", err)
	}
	println(createRequest.GetId())

	_, err = client.Update(ctx, &desc.UpdateRequest{Id: 12,
		Info: &desc.UpdateNoteInfo{Name: wrapperspb.String("Ivan"),
			Email: wrapperspb.String("example@example.com")}})
	if err != nil {
		log.Println("failed to update node", err)
	} else {
		fmt.Println("update request completed successfully")
	}

	_, err = client.Delete(ctx, &desc.DeleteRequest{Id: 12})
	if err != nil {
		log.Println("failed to delete node", err)
	} else {
		fmt.Println("Node deleted successfully")
	}

}
