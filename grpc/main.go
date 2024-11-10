package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	pb "stock_go/protos"
	"time"

	"google.golang.org/grpc"
)

type StockServer struct {
	pb.UnimplementedStockServiceServer
}

func (s *StockServer) StreamStockPrice(req *pb.StockRequest, stream pb.StockService_StreamStockPriceServer) error {
	price := 150.00
	for {
		priceChange := (rand.Float64() - 0.5) * 2.5
		price = price + priceChange

		res := &pb.StockResponse{
			StockSymbol: req.StockSymbol,
			Price:       price,
			Change:      priceChange,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error while create grpc server: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStockServiceServer(grpcServer, &StockServer{})

	fmt.Println("grpc Server listening on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
}
