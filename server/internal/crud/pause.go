package crud

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "playlist/proto"
)

func (s *Server) Pause(context.Context, *emptypb.Empty) (*pb.PauseStatus, error) {
	return &pb.PauseStatus{Error: "i am play"}, nil
}
