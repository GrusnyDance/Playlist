package crud

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
	pb "playlist/proto"
	"playlist/server/entity"
)

const chunkSize = 4096

func (s *Server) Play(n *emptypb.Empty, svr pb.Playlist_PlayServer) error {
	fmt.Println("hello")
	tr := &entity.Track{
		Name:          "Три дня дождя — Перезаряжай",
		Duration:      180,
		CurrentOffset: 0,
		Next:          nil,
		Prev:          nil,
	}
	s.PlayList.Tracks = tr
	s.PlayList.CurrentCursor = tr
	s.PlayList.LastTrack = tr
	s.PlayList.NumOfTracks = 1

	if s.PlayList.IsPlayed && (s.PlayList.CurrentPlay.Name == s.PlayList.CurrentCursor.Name) {
		return fmt.Errorf("i am played")
	}

	if s.PlayList.NumOfTracks == 0 {
		return fmt.Errorf("no tracks to play")
	}

	curr := s.PlayList.CurrentCursor
	fmt.Println((*curr).Name)
	for curr != nil {
		// Open the MP3 file
		fmt.Println("I see the loop")
		f, err := os.Open("./tracks/" + curr.Name + ".mp3")
		fmt.Println(curr.Name)
		if err != nil {
			f.Close()
			continue
		}

		// Create MP3 decoder
		mp3Decoder, err := mp3.NewDecoder(f)
		fmt.Println("len is", mp3Decoder.Length())
		if err != nil {
			f.Close()
			continue
		}

		// Read and transport audio in chunks
		buf := make([]byte, chunkSize)
		for {
			//fmt.Println("I see internal loop")
			n, err := mp3Decoder.Read(buf)
			fmt.Println(buf)
			if err == io.EOF {
				break
			}
			if n == 0 {
				break
			}
			//fmt.Println(curr.Name)
			//fmt.Println(buf)
			svr.Send(&pb.Audio{AudioChunk: buf, ChunkSize: int32(n)})
			offset, _ := mp3Decoder.Seek(0, io.SeekCurrent)
			curr.CurrentOffset = offset
		}
		f.Close()
		curr.CurrentOffset = 0
		curr = curr.Next
	}
	return nil
}
