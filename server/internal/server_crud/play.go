package server_crud

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
	pb "playlist/proto"
)

const chunkSize = 4096

func (s *Server) Play(n *emptypb.Empty, svr pb.Playlist_PlayServer) error {
	if s.PlayList.NumOfTracks == 0 {
		return fmt.Errorf("no tracks to play")
	}

	if s.PlayList.IsPlayed && (s.PlayList.CurrentPlay.Name == s.PlayList.CurrentCursor.Name) {
		return fmt.Errorf("i am played")
	}
	s.PlayList.IsPlayed = true

	curr := s.PlayList.CurrentCursor
	s.PlayList.CurrentPlay = curr
OuterLoop:
	for curr != nil {
		if !s.PlayList.IsPlayed {
			break OuterLoop
		}

		// Open the MP3 file
		f, err := os.Open("./tracks/" + curr.Name + ".mp3")
		if err != nil {
			f.Close()
			continue
		}

		// Create MP3 decoder
		mp3Decoder, err := mp3.NewDecoder(f)
		if err != nil {
			f.Close()
			continue
		}
		mp3Decoder.Seek(curr.CurrentOffset, 0)

		// Read and transport audio in chunks
		buf := make([]byte, chunkSize)
		for {
			if !s.PlayList.IsPlayed {
				offset, _ := mp3Decoder.Seek(0, io.SeekCurrent)
				curr.CurrentOffset = offset
				break OuterLoop
			}
			n1, err1 := mp3Decoder.Read(buf)
			if err1 == io.EOF {
				fmt.Println("EOF")
				break
			}
			if n1 == 0 {
				fmt.Println("n1 0")
				break
			}
			svr.Send(&pb.Audio{AudioChunk: buf, ChunkSize: int32(n1)})
		}
		f.Close()
		curr.CurrentOffset = 0
		curr = curr.Next
		s.PlayList.CurrentPlay = curr
	}
	return nil
}
