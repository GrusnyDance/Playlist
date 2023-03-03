package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
	"os/signal"
	pb "playlist/proto"
	"playlist/server/internal/playlist_controller"
	"playlist/server/internal/server_crud"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		grpclog.Fatal(err)
	}
	flag.Parse()

	listener, err := net.Listen("tcp", ":9999")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	svr, err := server_crud.NewServer()
	if err != nil {
		grpclog.Fatal(err)
	}
	fmt.Println("hello here")

	if err = playlist_controller.Start(svr); err != nil {
		log.Fatal(err)
	}

	pb.RegisterPlaylistServer(grpcServer, svr)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		playlist_controller.Finish(svr)
		svr.DbInstance.Db.Close()
		os.Exit(0)
	}()

	grpcServer.Serve(listener)
}
