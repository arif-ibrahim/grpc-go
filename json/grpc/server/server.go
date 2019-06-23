package main

import (
	"context"
	"encoding/json"
	pb "github.com/arif/grpc-go/json/grpc/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
}

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Price  int
	Fruits []string
}

type Combined struct {
	Page   int
	Price  int
	Fruits []string
}

func (s *server) Playjson(ctx context.Context, in *pb.JsonRequest) (*pb.JsonReply, error) {
	res1 := Response1{}
	res2 := Response2{}
	json.Unmarshal(in.D1, &res1)
	json.Unmarshal(in.D2, &res2)
	comb := Combined{}
	comb.Page = res1.Page
	comb.Price = res2.Price
	comb.Fruits = append(res1.Fruits, res2.Fruits...)
	byt, _ := json.Marshal(comb)
	return &pb.JsonReply{COMB: byt}, nil
}

func main() {
	listnr, _ := net.Listen("tcp", port)
	srvr := grpc.NewServer()

	pb.RegisterJsonconvServer(srvr, &server{})

	err := srvr.Serve(listnr)
	if err != nil {
		log.Fatal("failed to serve...\n")
	}
}
