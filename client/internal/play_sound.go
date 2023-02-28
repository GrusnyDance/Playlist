package internal

import (
	"fmt"
	"github.com/hajimehoshi/oto"
	"io"
	"log"
	pb "playlist/proto"
)

const chunkSize = 4096

func PlaySound(stream *pb.Playlist_PlayClient) {
	otoCtx, err := oto.NewContext(44100, 2, 2, chunkSize)
	outputDevice := otoCtx.NewPlayer()
	if err != nil {
		fmt.Println(err)
		return
	}
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
				log.Fatalf("cannot receive %v", err)
			}
			if _, err2 = outputDevice.Write(resp.AudioChunk[:resp.ChunkSize]); err != nil {
				log.Fatalf("cannot play %v", err)
			}
		}
	}()
	<-done
}
