package crud

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/levigross/grequests"
	"google.golang.org/api/youtube/v3"
	"net/http"
	"os"
	pb "playlist/proto"
	"playlist/server/entity"
	"playlist/server/internal/duration"
)

func (s *Server) AddSong(ctx context.Context, in *pb.SongName) (*pb.AddStatus, error) {
	track := &entity.Track{
		Next: nil,
		Prev: nil,
	}

	searchResult, title, err := searchVideo(ctx, s.YoutubeService, in.Name)
	dur, err := duration.GetDuration(searchResult.Id.VideoId)
	if err != nil {
		return &pb.AddStatus{Error: "unable to get duration"}, err
	}
	track.Duration = dur
	track.Name = title

	link, err := getAudioLink(searchResult.Id.VideoId)
	if err != nil {
		return &pb.AddStatus{Error: "unable to get audio from youtube"}, err
	}
	err = downloadAudio(link, title)
	if err != nil {
		return &pb.AddStatus{Error: "unable to download audiofile"}, err
	}

	s.PlayList.NumOfTracks++
	if s.PlayList.NumOfTracks > 1 {
		track.Prev = s.PlayList.LastTrack
		s.PlayList.LastTrack.Next = track
	} else {
		s.PlayList.CurrentCursor = track
	}
	s.PlayList.LastTrack = track

	return &pb.AddStatus{Error: "no errors", NewSongName: track.Name}, nil
}

func downloadAudio(link, title string) error {
	fmt.Println("I download")

	// Create the HTTP request
	resp, err := grequests.Get(link, nil)
	if err != nil {
		return err
	}

	fmt.Println("I finished get")
	if err = resp.DownloadToFile("./tracks/" + title + ".mp3"); err != nil {
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
