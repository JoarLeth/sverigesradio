package program

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/joarleth/sverigesradio/station"
	"io/ioutil"
	"net/http"
	"strings"
)

type Root struct {
	XMLName     xml.Name  `xml:"sr"`
	ProgramList []Program `xml:"programs>program"`
}

type Program struct {
	ExternalId  int    `xml:"id,attr"`
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Image       string `xml:"socialimage"`
}

func GetProgram(programName string, stationName string) (Program, error) {
	programList, _ := GetPrograms(stationName)

	for _, program := range programList {
		if strings.ToLower(program.Name) == strings.ToLower(programName) {
			return program, nil
		}
	}

	return Program{}, errors.New(fmt.Sprintf("Could not find program with name %s", programName))
}

func GetPrograms(stationName string) ([]Program, error) {
	station, _ := station.GetStation("P3")

	programsUrl := fmt.Sprintf("http://api.sr.se/api/v2/programs/index?pagination=false&isarchived=false&channelid=%d", station.ExternalId)

	resp, httpErr := http.Get(programsUrl)
	defer resp.Body.Close()

	if httpErr != nil {
		return []Program{}, errors.New("Get request failed in GetPrograms.")
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotModified) {
		return []Program{}, errors.New(fmt.Sprintf("GET request in GetPrograms returned status %d rather than %d or %d", resp.StatusCode, http.StatusOK, http.StatusNotModified))
	}

	// Only returns error if bytes Buffer becomes too large. How can this be tested?
	body, ioutilErr := ioutil.ReadAll(resp.Body)

	if ioutilErr != nil {
		return []Program{}, errors.New("ioutil.ReadAll failed in GetPrograms.")
	}

	program, _ := ExtractProgramsFromXML(body)

	return program, nil
}

func ExtractProgramsFromXML(xml_data []byte) ([]Program, error) {
	var r Root

	err := xml.Unmarshal(xml_data, &r)

	if err != nil {
		return nil, errors.New("program: unable to unmarshal xml_data in ExtractProgramsFromXML")
	}

	return r.ProgramList, nil
}
