package episode

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestExtractEpisodesFromXMLReturnsCorrectNumberOfEpisodes(t *testing.T) {
	xmlFile, _ := os.Open("sr_p1_vh_episodes.xml")

	defer xmlFile.Close()

	xml_data, _ := ioutil.ReadAll(xmlFile)

	episode_list, _ := ExtractEpisodesFromXML(xml_data)

	if len(episode_list) != 10 {
		t.Errorf("Unexpected number of episodes. Expected 10, got %d", len(episode_list))
	}
}

func TestExtractEpisodesFromXMLFirstEpisodeCorrect(t *testing.T) {
	xmlFile, _ := os.Open("sr_p1_vh_episodes.xml")

	defer xmlFile.Close()

	xml_data, _ := ioutil.ReadAll(xmlFile)

	episode_list, _ := ExtractEpisodesFromXML(xml_data)

	expected := Episode{
		ExternalId:   298577,
		Title:        "Kvinnokamp för skilsmässor och myndighet",
		Description:  "Vid sekelskiftet 1800 stod kvinnokampen om skilsmässor och rätten att bli sedd som en myndig person. Vetenskapsradion Historia träffar två forskare som undersökt hur ett fåtal kvinnor gick i bräschen för emancipationen i Sverige. – När det är maskerad kan jag inte hålla på mig, ska adelsdamen Aurora de Geer ha sagt om sitt frisläppta leverne. Historikern Carin Bergström har undersökt hur Aurora och andra högadelsdamer letade efter kryphål i lagen för att lyckas skilja sig med sina män. Bland annat var det vanligt att man hittade på otrohetsaffärer för att övertyga domstolarna om att godkänna en skilsmässa. Under 1800-talet var dessutom de flesta kvinnor omyndiga, och endast genom att söka dispens hos kungen kunde en ogift kvinna få tillgång till sitt arv. Historikern Britt Liljewall har undersökt de kvinnor som vågade trotsa normerna och ta strid för kvinnomyndigheten. – Med tiden kom frågan om myndighet att handla mindre om ekonomi och mer om identitet, säger hon. Programledare är Tobias Svanelid. Detta program är en repris från den 10 oktober 2013.",
		PublishedUTC: "2014-01-02T12:35:00Z",
		Duration:     1495,
	}
	actual := episode_list[0]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting episode not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestExtractEpisodesFromXMLErrorIsNil(t *testing.T) {
	xmlFile, _ := os.Open("sr_p1_vh_episodes.xml")

	defer xmlFile.Close()

	xml_data, _ := ioutil.ReadAll(xmlFile)

	_, err := ExtractEpisodesFromXML(xml_data)

	if err != nil {
		t.Errorf("Expected error to be nil. Got: %v", err.Error())
	}
}

func TestExtractEpisodesFromXMLInvalidXMLCausesError(t *testing.T) {
	xml_data := `
	<sr>
	  <copyright>
	</sr>
	`

	_, err := ExtractEpisodesFromXML([]byte(xml_data))

	if err == nil {
		t.Error("Expexted error caused by malformed xml.")
	}
}

func TestExtractEpisodesFromXMLInvalidXMLErrorMessage(t *testing.T) {
	xml_data := `
	<sr>
	  <copyright>
	</sr>
	`

	_, err := ExtractEpisodesFromXML([]byte(xml_data))

	expected := "episode: unable to unmarshal xml_data in ExtractEpisodesFromXML"
	actual := err.Error()

	if expected != actual {
		t.Errorf("Unexpected error message.\nExpected: %v\nActual: %v", expected, actual)
	}
}

// TODO: This will change. Use dependency injection.
func TestGetEpisodes(t *testing.T) {
	expected := Episode{
		ExternalId:   400445,
		Title:        "DIN PLAYLIST: Nadine Appelqvist med låtarna som ligger närmst hjärtat",
		Description:  "I veckans upplaga av Din Playlist är det Nadine Appelqvist som plockat ihop en timme musik. Nadine är även känd som vinnare av Musikguiden i P3:s fanquiz om Kent tidigare i somras.",
		PublishedUTC: "2014-07-14T20:03:00Z",
		Duration:     3420,
	}
	episodeList, _ := GetEpisides("Musikguiden i P3 ", "P3")

	actual := episodeList[0]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting episode not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}
