package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-golang/heroacademy-proto"
	"log"
	"net"
)

type server struct {
	heroacademy_proto.UnimplementedTrainingServiceServer
}

func (s *server) TrainHero(ctx context.Context, req *heroacademy_proto.HeroRequest) (*heroacademy_proto.HeroResponse, error) {
	heroName := req.GetHeroName()
	quirk := req.GetQuirk()

	result := fmt.Sprintf("Hero %s trained with quirk %s! Ready to save the day!", heroName, quirk)

	return &heroacademy_proto.HeroResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	heroacademy_proto.RegisterTrainingServiceServer(s, &server{})

	fmt.Printf("Serving on localhost:50051 ...\n")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
