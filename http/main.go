package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "stock_go/protos"

	"google.golang.org/grpc"
)

type StockHandler struct {
	gRPClient pb.StockServiceClient
}

func (s *StockHandler) stockUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	stream, err := s.gRPClient.StreamStockPrice(context.Background(), &pb.StockRequest{StockSymbol: "AAPL"})
	if err != nil {
		http.Error(w, "Failed to stream stock price", http.StatusInternalServerError)
		return
	}

	log.Println("Receiving stocks updates...")

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Println("Stream ended:", err)
			break
		}

		data, err := json.Marshal(res)

		if err != nil {
			log.Printf("Error marshalling update to JSON: %v", err)
			return
		}

		fmt.Fprintf(w, "data: %s\n\n", data)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewStockServiceClient(conn)

	stockHandler := &StockHandler{
		gRPClient: client,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/index.html")
	})
	http.HandleFunc("/stock-update", stockHandler.stockUpdateHandler)

	log.Println("http server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
