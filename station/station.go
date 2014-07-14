package station

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	SiteURL    string `xml:"siteurl"`
	LiveAudio  string `xml:"liveaudio>url"`
}

func GetStations() ([]Station, error) {
	resp, httpErr := http.Get("http://api.sr.se/api/v2/channels/index")
	defer resp.Body.Close()

	if httpErr != nil {
		return []Station{}, errors.New("Get request failed in GetStations.")
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotModified) {
		return []Station{}, errors.New(fmt.Sprintf("GET request in GetStations returned status %d rather than %d or %d", resp.StatusCode, http.StatusOK, http.StatusNotModified))
	}

	// Only returns error if bytes Buffer becomes too large. How can this be tested?
	body, ioutilErr := ioutil.ReadAll(resp.Body)

	if ioutilErr != nil {
		return []Station{}, errors.New("ioutil.ReadAll failed in GetStations.")
	}

	stations, _ := extractStationsFromXML(body)

	return stations, nil
}

func GetStation(name string) (Station, error) {
	stationList, _ := GetStations()

	for _, station := range stationList {
		if strings.ToLower(station.Name) == strings.ToLower(name) {
			return station, nil
		}
	}

	return Station{}, errors.New(fmt.Sprintf("Could not find station with name %s", name))
}

func extractStationsFromXML(xml_data []byte) ([]Station, error) {
	var r Root

	err := xml.Unmarshal(xml_data, &r)

	if err != nil {
		return nil, errors.New("station: unable to unmarshal xml_data in extractStationsFromXML")
	}

	return r.StationList, nil
}
