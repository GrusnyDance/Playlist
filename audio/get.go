package audio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	_ "github.com/silbinarywolf/preferdiscretegpu"
	"github.com/zergon321/reisen"
)

const (
	sampleRate                        = 44100
	channelCount                      = 2
	bitDepth                          = 8
	sampleBufferSize                  = 32 * channelCount * bitDepth * 1024
	SpeakerSampleRate beep.SampleRate = 44100
)

// readVideoAndAudio reads audio frames
// from the opened media and sends the decoded
// data to che channels to be played.
func readAudio(media *reisen.Media, sampleBuffer chan [2]float64) {
	err := media.OpenDecode()
	dTest, _ := media.Duration()
	fmt.Println("media duration is", dTest)

	if err != nil {
		fmt.Println(err, "34")
	}

	audioStream := media.AudioStreams()[0]
	fmt.Println(audioStream.BitRate(), "line 36")
	err = audioStream.Open()

	if err != nil {
		fmt.Println(err, "39")
	}

	fmt.Println("hello 42")

	for {
		packet, gotPacket, err := media.ReadPacket()
		packet, _, _ = media.ReadPacket()
		packet, _, _ = media.ReadPacket()

		if !gotPacket {
			fmt.Println("no packet")
			break
		}

		//fmt.Println(packet.Data(), "packet data is ")

		/*hash := sha256.Sum256(packet.Data())
		fmt.Println(base58.Encode(hash[:]))*/
		s := media.Streams()[packet.StreamIndex()].(*reisen.AudioStream)
		fmt.Println(s.FrameCount(), "frame count", packet.StreamIndex(), "stream index")
		audioFrame, gotFrame, err := s.ReadAudioFrame()
		//fmt.Println(audioFrame.Data(), "data is data")

		if !gotFrame {
			fmt.Println("no frame")
			break
		}
		if audioFrame == nil {
			fmt.Println("frame nil")
			continue
		}

		// Turn the raw byte data into
		// audio samples of type [2]float64.
		reader := bytes.NewReader(audioFrame.Data())
		//fmt.Println("i am data", audioFrame.Data())

		// See the README.md file for
		// detailed scheme of the sample structure.
		for reader.Len() > 0 {
			sample := [2]float64{0, 0}
			var result float64
			err = binary.Read(reader, binary.LittleEndian, &result)

			if err != nil {
				fmt.Println("error line 74")
			}
			//fmt.Println("result line 79", result)
			sample[0] = result

			err = binary.Read(reader, binary.LittleEndian, &result)

			if err != nil {
				fmt.Println("error line 82")
			}
			//fmt.Println("result line 87", result)

			sample[1] = result
			sampleBuffer <- sample
		}

		audioStream.Close()
		media.CloseDecode()
		close(sampleBuffer)
	}
}

// streamSamples creates a new custom streamer for
// playing audio samples provided by the source channel.
//
// See https://github.com/faiface/beep/wiki/Making-own-streamers
// for reference.
func streamSamples(sampleSource <-chan [2]float64) beep.Streamer {
	fmt.Println("hello 105")
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		numRead := 0

		fmt.Println("hello 109")

		for i := 0; i < len(samples); i++ {
			sample, ok := <-sampleSource

			if !ok {
				numRead = i + 1
				break
			}

			samples[i] = sample
			numRead++
		}

		if numRead < len(samples) {
			return numRead, false
		}
		fmt.Println("hello 124")
		return numRead, true
	})
}

// Strarts reading samples and frames
// of the media file.
func Start(filename string) error {
	fmt.Println("hello 143")
	// Initialize the audio speaker.
	err := speaker.Init(sampleRate,
		SpeakerSampleRate.N(time.Second/10))

	if err != nil {
		return fmt.Errorf("line 148 %v", err)
	}

	// Open the media file.
	media, err := reisen.NewMedia(filename)

	if err != nil {
		return fmt.Errorf("line 155 %v", err)
	}

	// Start decoding streams.
	sampleSource := make(chan [2]float64, sampleBufferSize)
	go readAudio(media, sampleSource)
	a := <-sampleSource
	fmt.Println(a, "i am a")
	b := <-sampleSource
	fmt.Println(b, "i am b")

	if err != nil {
		return fmt.Errorf("line 163 %v", err)
	}

	// Start playing audio samples.
	speaker.Play(streamSamples(sampleSource))
	fmt.Println("hello 169")

	return nil
}

func Get(filename string) {
	//track := &PlayerTest{}
	err := Start(filename + ".mp3")
	if err != nil {
		fmt.Println(err)
	}
}
