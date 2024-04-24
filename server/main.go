package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/hrojas2021/grpc/proto"
)

// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

// Ensure that calculatorServer implements CalculatorServer
// by embedding UnimplementedCalculatorServer
type calculatorServer struct {
	proto.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	result := req.A + req.B
	log.Printf("Receive new Add operation: %d + %d = %d", req.A, req.B, result)
	return &proto.AddResponse{Result: result}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, req *proto.SubtractRequest) (*proto.SubtractResponse, error) {
	result := req.A - req.B
	log.Printf("Receive new Subtract operation: %d - %d = %d", req.A, req.B, result)
	return &proto.SubtractResponse{Result: result}, nil
}

func (s *calculatorServer) Complex(ctx context.Context, req *proto.ComplexRequest) (*proto.ComplexResponse, error) {
	fmt.Printf("Request Received: %+v", req)

	var data [][]byte
	for _, h := range req.Hobbies {
		data = append(data, []byte(h))
	}

	response := &proto.ComplexResponse{
		Message:  "OK",
		Value:    345.45,
		Data:     data,
		IdToName: map[int32]string{200: "OK"},
		NestedMessage: &proto.NestedMessage{
			Timestamp: int64(300 * time.Microsecond),
			Count:     &wrapperspb.Int32Value{Value: 115},
		},
		NullableString: wrapperspb.String("Response back to client"),
		NullableInt32:  nil,
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterCalculatorServer(server, &calculatorServer{})

	log.Println("Starting gRPC server on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
