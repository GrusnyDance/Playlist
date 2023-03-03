package usecase

import (
	"context"
	"github.com/hajimehoshi/oto"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"playlist/client/internal"
	pb "playlist/proto"
	"time"
)

func New(client *pb.PlaylistClient, otoCtx *oto.Context) {
	//for i := 0; i < 5; i++ {
	//	go func(j int) {
	//		sn := &pb.SongName{
	//			Name: "валерий меладзе в первый день весны",
	//		}
	//		response, err := (*client).AddSong(context.Background(), sn)
	//		if err != nil {
	//			grpclog.Errorf("fail to add: %v", response)
	//		}
	//	}(i)
	//}

	time.Sleep(time.Second * 10)

	sn := &pb.SongName{
		Name: "Валерий Меладзе - Иностранец",
	}
	//response, err := (*client).AddSong(context.Background(), sn)
	//if err != nil {
	//	grpclog.Errorf("fail to add: %v", response)
	//}

	empty := &emptypb.Empty{}
	stream, err := (*client).Play(context.Background(), empty)
	if err != nil {
		grpclog.Errorf("error while playing %v", err)
	} else {
		go internal.PlaySound(&stream, otoCtx)
	}

	time.Sleep(time.Second * 15)
	(*client).Pause(context.Background(), empty)
	time.Sleep(time.Second * 5)
	stream, err = (*client).Play(context.Background(), empty)
	if err != nil {
		grpclog.Errorf("error while playing %v", err)
	} else {
		go internal.PlaySound(&stream, otoCtx)
	}

	time.Sleep(time.Second * 20)
	(*client).DeleteSong(context.Background(), sn)
}
