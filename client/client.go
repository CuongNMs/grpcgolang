package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcgolang/calculatorpb"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer cc.Close()
	client := calculatorpb.NewCalculatorServiceClient(cc)
	log.Println(client)
	callSum(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("callSum calling...")
	response, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 10,
		Num2: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.GetResult())
}
