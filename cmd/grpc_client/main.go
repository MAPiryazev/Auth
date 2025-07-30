package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

}
