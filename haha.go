package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
	"playlist/audio"
	"playlist/server/internal/duration"
)

func haha() {
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

	searchResult, title, err := searchVideo(ctx, youtubeService, "перезаряжай")
	dur, err := duration.GetDuration(searchResult.Id.VideoId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dur)

	link, err := getAudioLink(searchResult.Id.VideoId)
	if err != nil {
		fmt.Println(err)
	}
	err = downloadAudio(link, title)
	if err != nil {
		fmt.Println(err)
	}

	audio.Get(title + ".mp3")
}

func downloadAudio(link, title string) error {
	fmt.Println("I download")

	// Create the HTTP request
	resp, err := grequests.Get(link, nil)
	if err != nil {
		return err
	}

	fmt.Println("I finished get")
	if err = resp.DownloadToFile("./" + title + ".mp3"); err != nil {
		return fmt.Errorf("unable to download file: %v", err)
	}
	fmt.Println("I finished download to file")
	defer resp.ClearInternalBuffer()

	fmt.Println("Downloaded")
	return nil
}

func searchVideo(ctx context.Context, youtubeService *youtube.Service, query string) (*youtube.SearchResult, string, error) {
	searchCall := youtubeService.Search.List("snippet").
		Q(query).
		Order("relevance").
		Order("viewCount").
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
	fmt.Println("link is", link)
	return link, nil
}
