package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/cavaliergopher/grab/v3"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
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

	searchResult, title, err := searchVideo(ctx, youtubeService, "lalala")
	link, size, err := getAudioLink(searchResult.Id.VideoId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("size is", size)
	err = downloadAudio(link, title)
	if err != nil {
		fmt.Println(err)
	}
}

func downloadAudio(link, title string) error {
	resp, err := grab.Get(".", link)
	os.Rename(resp.Filename, title+".mp3")
	return err
}

func searchVideo(ctx context.Context, youtubeService *youtube.Service, query string) (*youtube.SearchResult, string, error) {
	searchCall := youtubeService.Search.List([]string{"snippet", "id"}).
		Q(query).
		Type("video").
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

func getAudioLink(videoId string) (link, size string, err error) {
	url := os.Getenv("VEVIOZ_API") + videoId
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", "", err
	}

	// Parse
	str := doc.Find("a")
	str.Each(func(i int, s *goquery.Selection) {
		if i == 2 {
			size = s.Find("div.text-shadow-1").Next().Text()
			link, _ = s.Attr("href")
		}
	})
	return link, size, nil
}
