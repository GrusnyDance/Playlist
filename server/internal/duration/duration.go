package duration

import (
	"encoding/json"
	duration "github.com/channelmeter/iso8601duration"
	"github.com/levigross/grequests"
	"os"
)

type VideoDescription struct {
	Kind  string `json:"-"`
	Etag  string `json:"-"`
	Items []struct {
		Kind           string `json:"-"`
		Etag           string `json:"-"`
		Id             string `json:"-"`
		ContentDetails struct {
			Duration        string `json:"duration"`
			Dimension       string `json:"-"`
			Definition      string `json:"-"`
			Caption         string `json:"-"`
			LicensedContent bool   `json:"-"`
			ContentRating   struct {
			} `json:"-"`
			Projection string `json:"-"`
		} `json:"contentDetails"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"-"`
		ResultsPerPage int `json:"-"`
	} `json:"-"`
}

func GetDuration(videoId string) (int, error) {
	link := os.Getenv("GOOGLE_API") + "id=" + videoId + "&part=contentDetails&key=" + os.Getenv("API_KEY")
	resp, err := grequests.Get(link, nil)
	if err != nil {
		return 0, err
	}

	var videoTags VideoDescription
	err = json.NewDecoder(resp.RawResponse.Body).Decode(&videoTags)
	if err != nil {
		return 0, err
	}

	dur, _ := duration.FromString(videoTags.Items[0].ContentDetails.Duration)
	return dur.Minutes*60 + dur.Seconds, nil
}
