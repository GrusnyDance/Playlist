package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"github.com/joho/godotenv"
	"github.com/melbahja/got"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	searchResult, title, err := searchVideo(ctx, youtubeService, "мукка")
	link, err := getAudioLink(searchResult.Id.VideoId)
	if err != nil {
		fmt.Println(err)
	}
	err = downloadAudio(link, title)
	if err != nil {
		fmt.Println(err)
	}

	playAudio(title)
}

func downloadAudio(link, title string) error {
	fmt.Println("I download")
	g := got.New()
	var err error
	err = g.Download(link, "./"+title+".mp3")
	fmt.Println("I downloaded")
	return err
}

func searchVideo(ctx context.Context, youtubeService *youtube.Service, query string) (*youtube.SearchResult, string, error) {
	searchCall := youtubeService.Search.List([]string{"snippet", "id"}).
		Q(query).
		Order("relevance").
		Order("rating").
		Type("video").
		VideoDuration("short").
		MaxResults(1)
	searchResult, err := searchCall.Do()
	if err != nil {
		return nil, "", fmt.Errorf("failed to search for video: %v", err)
	}
	if len(searchResult.Items) == 0 {
		return nil, "", fmt.Errorf("no video found for query '%s'", query)
	}
	title := searchResult.Items[0].Snippet.Title
	return searchResult.Items[0], title, nil
}

func getAudioLink(videoId string) (link string, err error) {
	url := os.Getenv("VEVIOZ_API") + videoId
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	// Parse
	str := doc.Find("a")
	str.Each(func(i int, s *goquery.Selection) {
		if i == 2 {
			link, _ = s.Attr("href")
		}
	})
	return link, nil
}

func playAudio(filename string) {
	data, e1 := os.Open(filename + ".mp3")
	defer data.Close()

	if e1 != nil {
		log.Fatalln(e1.Error())
	}
	decodedStream, e2 := mp3.NewDecoder(data)
	if e2 != nil {
		log.Fatalln(e2.Error())
	}
	otoCtx, readyChan, e3 := oto.NewContext(44100, 2, 2)
	if e3 != nil {
		log.Fatalln(e3.Error())
	}
	//ждем завершения инициализации
	<-readyChan
	player := otoCtx.NewPlayer(decodedStream)

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	player.Play()

	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
	// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
	// if err != nil{
	//     panic("player.Seek failed: " + err.Error())
	// }
	// println("Player is now at position:", newPos)
	// player.Play()

	// If you don't want the player/sound anymore simply close
	err := player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
