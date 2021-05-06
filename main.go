package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// type ProcessDefinition struct {
// }

type MainForm struct {
	FormID   string `xml:"FormID"`
	FormName string `xml:"FormName"`
	First    string `xml:"first"`
	Second   string `xml:"second"`
	Third    string `xml:"third"`
}

type DataSet struct {
	// ProcessDefinition `xml:""`
	MainForm `xml:"MainForm"`
}

type Person struct {
	XMLName    xml.Name `xml:"person"`
	FirstName  string   `xml:"firstName"`
	MiddleName string   `xml:"middleName"`
	LastName   string   `xml:"lastName"`
	Age        int64    `xml:"age"`
	Skills     []Skill  `xml:"skills"`
}

type Skill struct {
	XMLName        xml.Name `xml:"skill"`
	Name           string   `xml:"skillName"`
	YearsPracticed int64    `xml:"practice"`
}

func main() {

	d := []Skill{}
	sk := Skill{
		Name:           "aa",
		YearsPracticed: 11,
	}
	d = append(d, sk)
	v := Person{}
	v.Skills = d

	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
