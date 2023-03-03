package internal

import (
	"github.com/hajimehoshi/oto"
	_ "github.com/hajimehoshi/oto"
	"io"
	"log"
	pb "playlist/proto"
)

func PlaySound(stream *pb.Playlist_PlayClient, otoCtx *oto.Context) {
	outputDevice := otoCtx.NewPlayer()
	defer outputDevice.Close()
	done := make(chan bool)

	go func() {
		for {
			resp, err2 := (*stream).Recv()
			if err2 == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err2 != nil {
				log.Fatalf("cannot receive %v", err2)
			}
			if _, err2 = outputDevice.Write(resp.AudioChunk[:resp.ChunkSize]); err2 != nil {
				log.Fatalf("cannot play %v", err2)
			}
		}
	}()
	<-done
}
