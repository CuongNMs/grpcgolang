package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcgolang/calculatorpb"
	"log"
	"net"
)

type server struct {
	calculatorpb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("Server connecting...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum is running...")
	response := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}
	return response, nil
}
