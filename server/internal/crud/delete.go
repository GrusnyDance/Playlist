package crud

import (
	"context"
	"os"
	pb "playlist/proto"
)

func (s *Server) DeleteSong(ctx context.Context, in *pb.SongName) (*pb.DeleteStatus, error) {
	if s.PlayList.CurrentPlay.Name == in.Name && s.PlayList.IsPlayed {
		return &pb.DeleteStatus{Error: "played song cannot be deleted"}, nil
	}

	temp := s.PlayList.Tracks
	for temp != nil {
		if temp.Name == in.Name {
			if s.PlayList.CurrentCursor.Name == in.Name {
				if temp.Prev != nil {
					s.PlayList.CurrentCursor = temp.Prev
				} else {
					s.PlayList.CurrentCursor = temp.Next
				}
			}
			if temp.Prev == nil {
				s.PlayList.Tracks = temp.Next
			} else if temp.Next == nil {
				s.PlayList.LastTrack = temp.Prev
			} else {
				lastB4Delete := temp.Prev
				nextAfterDelete := temp.Next
				lastB4Delete.Next = nextAfterDelete
				nextAfterDelete.Prev = lastB4Delete
			}
			os.Remove("./tracks/" + temp.Name + ".mp3")
			s.PlayList.NumOfTracks--
			break
		} else {
			temp = temp.Next
		}
	}
	return &pb.DeleteStatus{Error: "am fine thanks"}, nil
}
