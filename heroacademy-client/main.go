package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-golang/heroacademy-proto"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTrainingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	resp, err := client.TrainHero(ctx, &pb.HeroRequest{HeroName: "Deku", Quirk: "One For All"})

	if err != nil {
		log.Fatalf("Error while calling TrainHero: %v", err)
	}

	fmt.Println("Response from Hero Training Service: ", resp.GetResult())
}
