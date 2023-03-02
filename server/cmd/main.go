package main

import (
	"flag"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	pb "playlist/proto"
	"playlist/server/internal/playlist_controller"
	"playlist/server/internal/server_crud"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		grpclog.Fatal(err)
	}
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:9999")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	svr, err := server_crud.NewServer()
	if err != nil {
		grpclog.Fatal(err)
	}
	defer svr.DbInstance.Db.Close()

	if err = playlist_controller.Start(svr); err != nil {
		log.Fatal(err)
	}
	defer playlist_controller.Finish(svr)

	pb.RegisterPlaylistServer(grpcServer, svr)
	grpcServer.Serve(listener)
}
