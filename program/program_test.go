package program

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestExtractProgramsFromXML(t *testing.T) {
	xmlFile, _ := os.Open("sr_p1_programs.xml")

	defer xmlFile.Close()

	xml_data, _ := ioutil.ReadAll(xmlFile)

	program_list, err := ExtractProgramsFromXML(xml_data)

	if err != nil {
		t.Errorf("Received unexpected error: %v", err.Error())
	}

	if len(program_list) != 80 {
		t.Errorf("Unexpected number of programs. Expected 80, got %d", len(program_list))
	}

	expected := Program{ExternalId: 407,
		Name:        "Vetenskapsradion Historia",
		Description: "Om allt från runristningsgillen till våldtäktspolitik, från 1700-talsskorpor till historiska dataspel. Vi är där historien är, snart sagt överallt, eftersom allt ju har sin egen historia.",
		Image:       "http://sverigesradio.se/diverse/appdata/isidor/images/news_images/407/2216423_512_512.jpg",
	}
	actual := program_list[0]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting program not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}

}

func TestExtractProgramsFromXMLInvalidXML(t *testing.T) {
	xml_data := `
	<sr>
	  <copyright>
	</sr>
	`

	programs, err := ExtractProgramsFromXML([]byte(xml_data))

	if programs != nil {
		t.Error("Expexted programs to be nil when passing malformed xml.")
	}

	if err == nil {
		t.Error("Expexted error caused by malformed xml.")
	} else {
		expected := "program: unable to unmarshal xml_data in ExtractProgramsFromXML"
		actual := err.Error()

		if expected != actual {
			t.Errorf("Unexpected error message.\nExpected: %v\nActual: %v", expected, actual)
		}

	}

}
