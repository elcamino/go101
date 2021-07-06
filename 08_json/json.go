package jayson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kr/pretty"
)

type SpotifyArtist struct {
	Name         string            `bson:"name"`
	ID           string            `bson:"id"`
	URI          string            `bson:"uri"`
	Endpoint     string            `bson:"endpoint"`
	ExternalURLs map[string]string `bson:"external_urls"`
}

type SpotifyTrack struct {
	Artists []SpotifyArtist `json:"artists" bson:"artists"`
	// A list of the countries in which the track can be played,
	// identified by their ISO 3166-1 alpha-2 codes.
	AvailableMarkets []string `json:"available_markets" bson:"available_markets"`
	// The disc number (usually 1 unless the album consists of more than one disc).
	DiscNumber int `json:"disc_number" bson:"disc_number"`
	// The length of the track, in milliseconds.
	Duration int `json:"duration_ms" bson:"duration_ms"`
	// Whether or not the track has explicit lyrics.
	// true => yes, it does; false => no, it does not.
	Explicit bool `json:"explicit" bson:"explicit"`
	// External URLs for this track.
	ExternalURLs map[string]string `json:"external_urls" bson:"external_urls"`
	ExternalIDs  map[string]string `bson:"external_ids"`
	// A link to the Web API endpoint providing full details for this track.
	Endpoint string `json:"href" bson:"href"`
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	// A URL to a 30 second preview (MP3) of the track.
	PreviewURL string `json:"preview_url" bson:"preview_url"`
	// The number of the track.  If an album has several
	// discs, the track number is the number on the specified
	// DiscNumber.
	TrackNumber int    `json:"track_number" bson:"track_number"`
	URI         string `json:"uri" bson:"url"`
	Popularity  int    `bson:"popularity"`
}

func Jayson() {
	jsonData, err := ioutil.ReadFile("data/test.json")
	if err != nil {
		log.Fatal(err)
	}

	var track SpotifyTrack
	err = json.Unmarshal(jsonData, &track)
	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(track)

	trackBytes, err := json.Marshal(track)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(trackBytes))

}
