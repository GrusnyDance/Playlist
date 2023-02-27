package crud

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "playlist/proto"
)

func (s *Server) Prev(context.Context, *emptypb.Empty) (*pb.PrevStatus, error) {
	return &pb.PrevStatus{Error: "am fine thanks"}, nil
}
