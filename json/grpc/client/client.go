package main

import (
	"context"
	"fmt"
	pb "github.com/arif/grpc-go/json/grpc/protobuf"
	"google.golang.org/grpc"
	"io/ioutil"
	"time"
)

const (
	address = "localhost:50051"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	Check(err)
	defer conn.Close()

	cl := pb.NewJsonconvClient(conn)

	dat1, err := ioutil.ReadFile("json/grpc/test-data/response1.json")
	Check(err)
	dat2, err := ioutil.ReadFile("json/grpc/test-data/response2.json")
	Check(err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := cl.Playjson(ctx, &pb.JsonRequest{D1: dat1, D2: dat2})

	fmt.Println(string(res.COMB))
}
