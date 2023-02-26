package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "playlist/proto"
	"playlist/server/internal/crud"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	svr := crud.NewServer()
	pb.RegisterPlaylistServer(grpcServer, svr)
	grpcServer.Serve(listener)
}
