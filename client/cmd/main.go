package main

import (
	"context"
	"fmt"
	"github.com/hajimehoshi/oto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"os/signal"
	"playlist/client/usecase"
	pb "playlist/proto"
	"syscall"
)

func main() {
	conn, err := grpc.Dial("localhost:9999",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlaylistClient(conn)
	otoCtx, _ := oto.NewContext(44100, 2, 2, 4096)
	usecase.New(&client, otoCtx)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	fmt.Println("i am graceful")
	empty := &emptypb.Empty{}
	client.Pause(context.Background(), empty)
}
