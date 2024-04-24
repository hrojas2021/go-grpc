package main

import (
	"context"
	"log"
	"time"

	"github.com/hrojas2021/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewCalculatorClient(conn)

	addResult, err := client.Add(context.Background(), &proto.AddRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("Add RPC failed: %v", err)
	}
	log.Printf("Add Result: %d\n", addResult.Result)

	subResult, err := client.Subtract(context.Background(), &proto.SubtractRequest{A: 30, B: 15})
	if err != nil {
		log.Fatalf("Subtract RPC failed: %v", err)
	}

	log.Printf("Subtract Result: %d\n", subResult.Result)

	complexReq := &proto.ComplexRequest{
		Name:      "Complex Request",
		Age:       34,
		Height:    171.5,
		IsStudent: false,
		Hobbies:   []string{"Soccer", "Poker"},
		Scores:    map[string]int32{"Bayern Munchen": 3, "Real Madrid": 1},
		NestedMessage: &proto.NestedMessage{
			Timestamp: int64(time.Second),
			Count:     wrapperspb.Int32(342), // &wrapperspb.Int32Value{Value: 34}
		},
	}

	complexResponse, err := client.Complex(context.Background(), complexReq)
	if err != nil {
		log.Fatalf("Complex RPC failed: %v", err)
	}
	log.Printf("Complex Result: %+v\n", complexResponse)
}
