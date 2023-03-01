package usecase

import (
	"context"
	"fmt"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"playlist/client/internal"
	pb "playlist/proto"
)

func New(client *pb.PlaylistClient) {
	for i := 0; i < 5; i++ {
		go func(j int) {
			sn := &pb.SongName{
				Name: "валерий меладзе в первый день весны",
			}
			response, err := (*client).AddSong(context.Background(), sn)
			if err != nil {
				grpclog.Fatalf("fail to add: %v", response)
			}
		}(i)
	}

	sn := &pb.SongName{
		Name: "перезаряжай",
	}
	response, err := (*client).AddSong(context.Background(), sn)
	if err != nil {
		grpclog.Fatalf("fail to add: %v", response)
	}

	empty := &emptypb.Empty{}
	stream, err := (*client).Play(context.Background(), empty)
	if err != nil {
		fmt.Println("error while playing", err)
	}
	go internal.PlaySound(&stream)

}
