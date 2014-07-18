package episode

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/joarleth/sverigesradio/program"
	"io/ioutil"
	"net/http"
)

type Root struct {
	XMLName     xml.Name  `xml:"sr"`
	EpisodeList []Episode `xml:"episodes>episode"`
}

type Episode struct {
	ExternalId        int    `xml:"id,attr"`
	ExternalProgramId int    `xml:"-"`
	Title             string `xml:"title"`
	Description       string `xml:"description"`
	PublishedUTC      string `xml:"publishdateutc"`
	Duration          int    `xml:"broadcast>playlist>duration"`
}

func GetEpisides(programName string, stationName string) ([]Episode, error) {
	program, _ := program.GetProgram(programName, stationName)

	episodesUrl := fmt.Sprintf("http://api.sr.se/api/v2/episodes/index?programid=%d", program.ExternalId)

	println(episodesUrl)

	resp, httpErr := http.Get(episodesUrl)
	defer resp.Body.Close()

	if httpErr != nil {
		return []Episode{}, errors.New("Get request failed in GetEpisides.")
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotModified) {
		return []Episode{}, errors.New(fmt.Sprintf("GET request in GetEpisides returned status %d rather than %d or %d", resp.StatusCode, http.StatusOK, http.StatusNotModified))
	}

	// Only returns error if bytes Buffer becomes too large. How can this be tested?
	body, ioutilErr := ioutil.ReadAll(resp.Body)

	if ioutilErr != nil {
		return []Episode{}, errors.New("ioutil.ReadAll failed in GetEpisides.")
	}
	episodeList, _ := ExtractEpisodesFromXML(body)

	for i := range episodeList {
		episodeList[i].ExternalProgramId = program.ExternalId
	}

	return episodeList, nil
}

func ExtractEpisodesFromXML(xml_data []byte) ([]Episode, error) {
	var r Root

	err := xml.Unmarshal(xml_data, &r)

	if err != nil {
		return nil, errors.New("episode: unable to unmarshal xml_data in ExtractEpisodesFromXML")
	}

	return r.EpisodeList, nil
}
