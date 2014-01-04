package episode

import (
	"encoding/xml"
	"errors"
)

type Root struct {
	XMLName     xml.Name  `xml:"sr"`
	EpisodeList []Episode `xml:"episodes>episode"`
}

type Episode struct {
	ExternalId   int    `xml:"id,attr"`
	Title        string `xml:"title"`
	Description  string `xml:"description"`
	PublishedUTC string `xml:"publishdateutc"`
	Duration     int    `xml:"broadcast>playlist>duration"`
}

func ExtractEpisodesFromXML(xml_data []byte) ([]Episode, error) {
	var r Root

	err := xml.Unmarshal(xml_data, &r)

	if err != nil {
		return nil, errors.New("episode: unable to unmarshal xml_data in ExtractEpisodesFromXML")
	}

	return r.EpisodeList, nil
}
