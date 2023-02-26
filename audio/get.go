package audio

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"os"
	"time"
)

func Get(filename string) {
	const chunkSize = 4096

	// Open the MP3 file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create MP3 decoder
	mp3Decoder, err := mp3.NewDecoder(f)
	if err != nil {
		panic(err)
	}

	// Open audio output device
	otoCtx, _ := oto.NewContext(44100, 2, 2, chunkSize)
	outputDevice := otoCtx.NewPlayer()
	if err != nil {
		panic(err)
	}
	defer outputDevice.Close()

	// Read and play audio in chunks
	buf := make([]byte, chunkSize)
	for {
		n, err := mp3Decoder.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if n == 0 {
			break
		}
		// Write audio data to output device
		if _, err := outputDevice.Write(buf[:n]); err != nil {
			panic(err)
		}

		// Sleep to simulate audio transport delay
		time.Sleep(5 * time.Millisecond)
		// Calculate the percentage of the file that has been played
		offset, _ := mp3Decoder.Seek(0, io.SeekCurrent)
		fmt.Println("current offset", offset)
	}
}
