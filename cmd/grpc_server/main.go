package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/MAPiryazev/Auth/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

// запрос на получение данных пользователя
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponce, error) {
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "request cancelled or timed out")
	default:
		// Контекст активен, продолжаем
	}

	log.Printf("Note id: %d", req.GetId())

	// Генерация времени и преобразование в protobuf формат
	created := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())
	updated := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())

	return &desc.GetResponce{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		Role:      desc.NoteINfo_Role(gofakeit.Number(0, 2)),
		CreatedAt: timestamppb.New(created),
		UpdatedAt: timestamppb.New(updated),
	}, nil
}

// запрос на создание пользователя
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponce, error) {
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "request cancelled or timed out")
	default:
	}

	log.Println("Creation request")
	fmt.Println("Email ", req.Info.Email)
	fmt.Println("Name ", req.Info.Name)
	fmt.Println("Password ", req.Info.Password)
	fmt.Println("PasswordConfirm ", req.Info.PasswordCoonfirm)

	return &desc.CreateResponce{Id: gofakeit.Int64()}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	select {
	case <-ctx.Done():
		return &emptypb.Empty{}, status.Error(codes.DeadlineExceeded, "request cancelled or timed out")
	default:
	}

	log.Printf("Update request %d", req.GetId())
	fmt.Println("Email ", req.Info.GetEmail())
	fmt.Println("Name ", req.Info.GetName())
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	select {
	case <-ctx.Done():
		return &emptypb.Empty{}, status.Error(codes.DeadlineExceeded, "request cancelled or timed out")
	default:
	}

	log.Println("Delete request")
	fmt.Printf("User id requested for deletion %d", req.GetId())
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(grpcPort))
	if err != nil {
		log.Fatalf("failed to listen api port %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
