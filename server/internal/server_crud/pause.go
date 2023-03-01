package server_crud

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "playlist/proto"
)

func (s *Server) Pause(context.Context, *emptypb.Empty) (*pb.PauseStatus, error) {
	s.PlayList.IsPlayed = false
	return &pb.PauseStatus{Error: "am fine"}, nil
}
