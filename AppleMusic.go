package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

// AppleMusicResult -- apple music search result
type AppleMusicResult struct {
	WrapperType            string  `json:"wrapperType"`
	Kind                   string  `json:"kind"`
	ArtistId               int     `json:"artistId"`
	CollectionId           int     `json:"collectionId"`
	TrackId                int     `json:"trackId"`
	ArtistName             string  `json:"artistName"`
	CollectionName         string  `json:"collectionName"`
	TrackName              string  `json:"trackName"`
	CollectionCensoredName string  `json:"collectionCensoredName"`
	TrackCensoredName      string  `json:"trackCensoredName"`
	CollectionArtistName   string  `json:"collectionArtistName"`
	ArtistViewUrl          string  `json:"artistViewUrl"`
	CollectionViewUrl      string  `json:"collectionViewUrl"`
	TrackViewUrl           string  `json:"trackViewUrl"`
	PreviewUrl             string  `json:"previewUrl"`
	ArtworkUrl30           string  `json:"artworkUrl30"`
	ArtworkUrl60           string  `json:"artworkUrl60"`
	ArtworkUrl100          string  `json:"artworkUrl100"`
	CollectionPrice        float32 `json:"collectionPrice"`
	TrackPrice             float32 `json:"trackPrice"`
	ReleaseDate            string  `json:"releaseDate"`
	CollectionExplicitness string  `json:"collectionExplicitness"`
	TrackExplicitness      string  `json:"trackExplicitness"`
	DiscCount              int     `json:"discCount"`
	DiscNumber             int     `json:"discNumber"`
	TrackCount             int     `json:"trackCount"`
	TrackNumber            int     `json:"trackNumber"`
	TrackTimeMillis        int     `json:"trackTimeMillis"`
	Country                string  `json:"country"`
	Currency               string  `json:"currency"`
	PrimaryGenreName       string  `json:"primaryGenreName"`
	IsStreamable           bool    `json:"isStreamable"`
}

// AppleMusicResults -- apple music search results
type AppleMusicResults struct {
	Results     []AppleMusicResult `json:"results"`
	ResultCount int                `json:"resultCount"`
}

// AppleMusicLimit -- apple music search result count limit
const AppleMusicLimit = 10

// AppleMusicURL -- apple music search api scheme
const AppleMusicURL = "https://itunes.apple.com/search?country=jp&entity=album&lang=en_us&term=%s&limit=%d"

// SearchAppleMusic -- search apple music for album
func SearchAppleMusic(keyword string) *AppleMusicResults {
	url := fmt.Sprintf(AppleMusicURL, keyword, AppleMusicLimit)
	reader := strings.NewReader("")
	req, _ := http.NewRequest("GET", url, reader)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	//fmt.Println(result)
	r := &AppleMusicResults{}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		if err := json.Unmarshal([]byte(bodyString), r); err != nil {
			panic(err)
		}
	}
	defer resp.Body.Close()
	return r
}

// ProcessSearchAppleMusicEvent -- entry for search apple music event
func ProcessSearchAppleMusicEvent(command string, channelID string) {
	keyword := strings.Replace(command, "/music ", "", -1)
	result := SearchAppleMusic(strings.Replace(keyword, " ", "+", -1))
	for _, r := range result.Results {
		single := fmt.Sprintf("Artist: %s\nAlbum: %s\nRelease Date: %s\nListen on Apple Music: %s", r.ArtistName, r.CollectionName, r.ReleaseDate, r.CollectionViewUrl)
		//rtm.SendMessage(rtm.NewOutgoingMessage(single, channel))

		// TODO
		api := slack.New(os.Getenv("SLACK_TOKEN"))
		params := slack.PostMessageParameters{AsUser: true, UnfurlLinks: true, UnfurlMedia: true}
		channelName, timestamp, err := api.PostMessage(channelID, single, params)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Message successfully sent to channel %s at %s", channelName, timestamp)
	}
}
