package main

import (
	"flag"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "playlist/proto"
	"playlist/server/internal/server_crud"
	"playlist/server/internal/storage_playlist_controller"
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
	storage_playlist_controller.LoadFromDbToPlaylist(svr)
	pb.RegisterPlaylistServer(grpcServer, svr)
	grpcServer.Serve(listener)
}
