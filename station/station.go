package station

import (
	"encoding/xml"
	"errors"
)

type Root struct {
	XMLName     xml.Name  `xml:"sr"`
	StationList []Station `xml:"channels>channel"`
}

type Station struct {
	ExternalId int    `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	Image      string `xml:"image"`
	Color      string `xml:"color"`
	Tagline    string `xml:"tagline"`
	SiteURL    string `xml:"siteurl"`
	LiveAudio  string `xml:"liveaudio>url"`
}

func ExtractStationsFromXML(xml_data []byte) ([]Station, error) {
	var r Root

	err := xml.Unmarshal(xml_data, &r)

	if err != nil {
		return nil, errors.New("station: unable to unmarshal xml_data in ExtractStationsFromXML")
	}

	return r.StationList, nil
}
