package lesson13

import "encoding/xml"

type Address struct {
	Country string `xml:"country,attr"`
	City    string `xml:"city,attr"`
}

type Person struct {
	XMLName   xml.Name  `xml:"person"`
	Name      string    `xml:"name"`
	IsAwesome bool      `xml:"awesome,attr,omitempty"`
	Residence *Address  `xml:"residence"`
	Friends   []*Person `xml:"friends>person"`
}
