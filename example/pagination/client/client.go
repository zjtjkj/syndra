package main

import (
	"context"
	"fmt"

	"github.com/zjtjkj/syndra/example/pagination/object"
	"google.golang.org/grpc"
)

func main() {
	listConn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	listClient := object.NewObjectServiceClient(listConn)
	resp, err := listClient.ListObject(context.Background(), &object.ListRequest{
		Index: 1,
		Size:  2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}