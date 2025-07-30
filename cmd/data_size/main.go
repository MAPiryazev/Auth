package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/proto"

	desc "github.com/MAPiryazev/Auth/pkg/note_v1"
)

func main() {
	session := &desc.NoteINfo{
		Name:  gofakeit.BeerName(),
		Email: gofakeit.Email(),
	}

	dataJson, error := json.Marshal(session)
	if error != nil {
		log.Println(error)
	}
	fmt.Printf("\n\ndataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, error := proto.Marshal(session)
	if error != nil {
		log.Println(error)
	}
	fmt.Printf("dataPb len %d byte \n%v\n", len(dataPb), dataPb)
}
