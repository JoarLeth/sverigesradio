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

	station_list, err := ExtractStationsFromXML(xml_data)

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
		Tagline:   "den talade kanalen",
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

	stations, err := ExtractStationsFromXML([]byte(xml_data))

	if stations != nil {
		t.Error("Expexted stations to be nil when passing malformed xml.")
	}

	if err == nil {
		t.Error("Expexted error caused by malformed xml.")
	} else {
		expected := "station: unable to unmarshal xml_data in ExtractStationsFromXML"
		actual := err.Error()

		if expected != actual {
			t.Errorf("Unexpected error message.\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}
