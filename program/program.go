package program

import (
	"encoding/xml"
	"errors"
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

func ExtractProgramsFromXML(xml_data []byte) ([]Program, error) {
	var r Root

	err := xml.Unmarshal(xml_data, &r)

	if err != nil {
		return nil, errors.New("program: unable to unmarshal xml_data in ExtractProgramsFromXML")
	}

	return r.ProgramList, nil
}
