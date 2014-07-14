package station

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestExtractStationsFromXML(t *testing.T) {
	xmlFile, _ := os.Open("sr_stations.xml")

	defer xmlFile.Close()

	xml_data, _ := ioutil.ReadAll(xmlFile)

	station_list, err := extractStationsFromXML(xml_data)

	if err != nil {
		t.Errorf("Received unexpected error: %v", err.Error())
	}

	if len(station_list) != 10 {
		t.Errorf("Unexpected number of stations. Expected 10, got %d", len(station_list))
	}

	expected := Station{ExternalId: 132,
		Name:      "P1",
		Image:     "http://sverigesradio.se/diverse/appdata/isidor/images/news_images/132/2186746_512_512.jpg",
		Color:     "31a1bd",
		SiteURL:   "http://sverigesradio.se/p1",
		LiveAudio: "http://sverigesradio.se/topsy/direkt/132.mp3",
	}
	actual := station_list[0]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting station not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestExtractStationsFromXMLInvalidXml(t *testing.T) {
	xml_data := `
	<sr>
	  <copyright>
	</sr>
	`

	stations, err := extractStationsFromXML([]byte(xml_data))

	if stations != nil {
		t.Error("Expexted stations to be nil when passing malformed xml.")
	}

	if err == nil {
		t.Error("Expexted error caused by malformed xml.")
	} else {
		expected := "station: unable to unmarshal xml_data in extractStationsFromXML"
		actual := err.Error()

		if expected != actual {
			t.Errorf("Unexpected error message.\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

func TestDownloadAndExtractStations(t *testing.T) {
	expected := Station{ExternalId: 132,
		Name:      "P1",
		Image:     "http://sverigesradio.se/sida/images/132/2186745_512_512.jpg?preset=api-default-square",
		Color:     "31a1bd",
		SiteURL:   "http://sverigesradio.se/p1",
		LiveAudio: "http://sverigesradio.se/topsy/direkt/132.mp3",
	}

	stationList, _ := GetStations()

	actual := stationList[0]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting station not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGetStation(t *testing.T) {
	expected := Station{ExternalId: 164,
		Name:      "P3",
		Image:     "http://sverigesradio.se/sida/images/164/2186756_512_512.jpg?preset=api-default-square",
		Color:     "19a972",
		SiteURL:   "http://sverigesradio.se/p3",
		LiveAudio: "http://sverigesradio.se/topsy/direkt/164.mp3",
	}

	actual, _ := GetStation("P3")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting station not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}
