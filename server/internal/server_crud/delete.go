package server_crud

import (
	"context"
	pb "playlist/proto"
)

func (s *Server) DeleteSong(ctx context.Context, in *pb.SongName) (*pb.DeleteStatus, error) {
	//err := os.Remove("./tracks/" + in.Name + ".mp3")
	s.PlayList.Delete(in.Name)
	//if err != nil {
	//	return nil, err
	//}
	s.DbInstance.Delete(in.Name)
	return &pb.DeleteStatus{Error: "am fine thanks"}, nil
}
