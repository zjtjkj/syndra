package main

import (
	"context"
	"net"

	"github.com/zjtjkj/syndra/example/pagination/object"
	"github.com/zjtjkj/syndra/utils/pagination"
	"google.golang.org/grpc"
)

var data = []*object.Object{
	{Name: "objetc_1"}, {Name: "objetc_2"}, {Name: "objetc_3"}, {Name: "objetc_4"},
}

type ListServer struct {
	object.UnimplementedObjectServiceServer
}

func (l *ListServer) ListObject(_ context.Context, req *object.ListRequest) (resp *object.ListRepsonse, err error) {
	res, page := pagination.Page(data, req.Index, req.Size)
	resp = &object.ListRepsonse{
		Objects: res,
		Index:   req.Index,
		Size:    req.Size,
		Pages:   page,
		Total:   uint32(len(data)),
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	object.RegisterObjectServiceServer(s, &ListServer{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
