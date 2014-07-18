package track

import (
	"github.com/joarleth/sverigesradio/episode"
	"testing"
)

func TestGetTracks(t *testing.T) {
	episodeList, _ := episode.GetEpisides("Musikguiden i P3 ", "P3")

	GetTracks(episodeList[0])
}
