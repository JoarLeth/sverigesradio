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

func TestGetPrograms(t *testing.T) {
	expected := Program{ExternalId: 1646,
		Name:        "P3 Nyheter",
		Description: "Berättar vad som händer i Sverige och världen och håller extra koll på det som berör unga människor.",
		Image:       "http://sverigesradio.se/sida/images/1646/3184974_512_512.jpg?preset=api-default-square",
	}

	programList, _ := GetPrograms("P3")

	actual := programList[0]

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting station not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGetProgram(t *testing.T) {
	expected := Program{ExternalId: 4067,
		Name:        "Musikguiden i P3 ",
		Description: "Vi spelar den bästa nya musiken, bred och smal, kända hits och välbevarade hemligheter. I Musikguiden i P3 får du träffa artisterna, höra de nya låtarna och se hela sammanhanget. Ungefär som i naturfilm, fast med musik. Typ.",
		Image:       "http://sverigesradio.se/sida/images/4067/2472633_512_512.jpg?preset=api-default-square",
	}

	actual, _ := GetProgram("Musikguiden i P3 ", "P3")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Resulting program not matching expected.\nExpected: %v\nActual: %v", expected, actual)
	}
}
