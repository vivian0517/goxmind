package goxmind

import (
	"encoding/xml"
)

type Meta struct {
	XMLName xml.Name `xml:"meta"`
	Author  string   `xml:"author"`
	Created string   `xml:"created"`
}

type Content struct {
	XMLName    xml.Name    `xml:"xmap-content"`
	XMLSheets  []*XMLSheet `xml:"sheet"`
	Modified   string      `xml:"modified-by,attr,omitempty"`
	Xmlns      string      `xml:"xmlns,attr"`
	Xmlnsfo    string      `xml:"xmlns:fo,attr"`
	Xmlnssvg   string      `xml:"xmlns:svg,attr"`
	Xmlnsxhtml string      `xml:"xmlns:xhtml,attr"`
	Xmlnsxlink string      `xml:"xmlns:xlink,attr"`
}

type XMLSheet struct {
	Topic Topic  `xml:"topic"`
	Title string `xml:"title"`
	Id    string `xml:"id,attr"`
}

type Topic struct {
	XMLName        xml.Name   `xml:"topic"`
	Structureclass string     `xml:"structure-class,attr,omitempty"`
	Id             string     `xml:"id,attr"`
	Title          string     `xml:"title"`
	Children       *Children  `xml:"children,omitempty"`
	MakerRefs      *MakerRefs `xml:"marker-refs,omitempty"`
	Notes          *Note      `xml:"notes,omitempty"`
	Href           string     `xml:"xlink:href,attr,omitempty"`
}

type Children struct {
	Topics Topics `xml:"topics"`
}

type Topics struct {
	XMLName xml.Name `xml:"topics"`
	Type    string   `xml:"type,attr"`
	Topic   []Topic  `xml:"topic"`
}

type MakerRefs struct {
	MakerRef []MakerRef `xml:"marker-ref,omitempty"`
}

type MakerRef struct {
	MakerId string `xml:"marker-id,attr,omitempty"`
}

type Note struct {
	Plain string `xml:"plain,omitempty"`
}

type Manifest struct {
	XMLName      xml.Name    `xml:"manifest"`
	PasswordHint string      `xml:"password-hint,attr"`
	Xmlns        string      `xml:"xmlns,attr"`
	FileEntries  []FileEntry `xml:"file-entry"`
}

type FileEntry struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}
