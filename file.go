package goxmind

import (
	"archive/zip"
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//go:embed styles/styles.xml
var styles []byte
var themeid = "3t9ojddlutknqfparq91tnb3ls"
var provider = "org.xmind.ui.map.unbalanced"

func createManifest() *Manifest {
	return &Manifest{
		PasswordHint: "",
		Xmlns:        "urn:xmind:xmap:xmlns:manifest:1.0",
		FileEntries: []FileEntry{
			{FullPath: "attachments", MediaType: ""},
			{FullPath: "content.xml", MediaType: "text/xml"},
			{FullPath: "META-INF/", MediaType: ""},
			{FullPath: "META-INF/manifest.xml", MediaType: "text/xml"},
			{FullPath: "styles.xml", MediaType: "text/xml"},
		},
	}
}

// Save saves the xmind file, pass in the filename as a parameter
func (xmind *Xmind) Save(filename string) error {
	if strings.HasSuffix(filename, ".xmind") {
		xmind.FileName = filename
	} else {
		xmind.FileName = filename + ".xmind"
	}
	zipFile, err := os.Create(xmind.FileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	metaFile, err := archive.Create("META-INF/manifest.xml")
	if err != nil {
		return err
	}

	meta := createManifest()
	metaXML, err := xml.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	metaFile.Write(metaXML)

	contentFile, err := archive.Create("content.xml")
	if err != nil {
		return err
	}

	content := MarshalContent(xmind)
	contentXML, err := xml.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	contentFile.Write(contentXML)

	stylesFile, err := archive.Create("styles.xml")
	if err != nil {
		return err
	}
	stylesFile.Write(styles)
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	// The generated xmind is saved in
	p := "The generated xmind is saved in " + path + string(os.PathSeparator) + xmind.FileName
	fmt.Println(p)
	return nil
}

func Load(filename string) (*Xmind, error) {
	if !strings.HasSuffix(filename, ".xmind") {
		filename += ".xmind"
	}

	zipFile, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()

	var jsoncontent []byte
	var xmlcontent []byte
	var content *Content
	var xmind *Xmind

	for _, file := range zipFile.File {
		if strings.HasSuffix(file.Name, "content.json") {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			jsoncontent, err = io.ReadAll(rc)
			if err != nil {
				return nil, err
			}
		} else if strings.HasSuffix(file.Name, "content.xml") {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			xmlcontent, err = io.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			// This file can not be opened normally, please do not modify and save, otherwise the contents will be permanently lost！
			if strings.Contains(string(xmlcontent), "This file can not be opened normally, please do not modify and save, otherwise the contents will be permanently lost！") {
				continue
			}
		}
	}

	switch {
	case jsoncontent != nil:
		content, err = prasejsoncontent(jsoncontent)
		if err != nil {
			// Error parsing json
			log.Fatalf("Error parsing json: %v", err)
			return nil, err
		}
		xmind = UnmarshalContent(content)
		xmind.FileName = filename
		return xmind, nil
	case xmlcontent != nil:
		content, err = prasexmlcontent(xmlcontent)
		if err != nil {
			// Error parsing XML
			log.Fatalf("Error parsing XML: %v", err)
			return nil, err
		}
		xmind = UnmarshalContent(content)
		xmind.FileName = filename
		return xmind, nil
	default:
		return nil, errors.New("not found content file, invalid xmind")
	}
}

func prasejsoncontent(data []byte) (*Content, error) {
	newJson := &NewJson{}
	err := json.Unmarshal(data, newJson)
	if err != nil {
		// Error parsing json
		log.Fatalf("Error parsing json: %v", err)
		return nil, err
	}
	// Parsing json
	log.Println("Parsing json:", newJson)
	content := convertNewJsonToXmind(newJson)

	return content, nil
}
func prasexmlcontent(data []byte) (*Content, error) {
	content := &Content{}
	err := xml.Unmarshal(data, content)
	if err != nil {
		// Error parsing XML
		log.Fatalf("Error parsing XML: %v", err)
		return nil, err
	}
	return content, nil
}

func UnmarshalContent(content *Content) *Xmind {
	xmind := &Xmind{}
	for i := 0; i < len(content.XMLSheets); i++ {
		canva := &Sheet{
			SheetTitle: content.XMLSheets[i].Title,
			Node:       convertTopicToNode(&content.XMLSheets[i].Topic),
		}
		xmind.Sheets = append(xmind.Sheets, canva)
	}
	return xmind
}

func convertTopicToNode(topic *Topic) Node {
	node := Node{
		NodeTitle: topic.Title,
	}

	// Extract icons (marker-id)
	if topic.MarkerRefs != nil && len(topic.MarkerRefs.MarkerRef) > 0 {
		node.Markers = make([]Markers, len(topic.MarkerRefs.MarkerRef))
		for i, markerRef := range topic.MarkerRefs.MarkerRef {
			node.Markers[i] = Markers{Marker: markerRef.MarkerId}
		}
	}

	// Extract notes
	if topic.Notes != nil {
		node.Notes = topic.Notes.Plain
	}

	// Extract hyperlink
	if topic.Href != "" {
		node.Href = topic.Href
	}

	// Recursively process child nodes
	if topic.Children != nil && topic.Children.Topics.Type == "attached" {
		for _, childTopic := range topic.Children.Topics.Topic {
			childNode := convertTopicToNode(&childTopic)
			node.Children = append(node.Children, &childNode)
		}
	}
	return node
}

func (x *Xmind) PrintJson() {
	jsonBytes, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}
func (x *Xmind) SaveJson(filename string) {
	jsonBytes, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("save ", filename, " success")
}

func PraseJsonSaveXmind(filename string) {
	txt, _ := os.ReadFile("jsondata.txt")

	var xminddata *Xmind
	err := json.Unmarshal(txt, &xminddata)
	if err != nil {
		fmt.Println(err)
	}
	xminddata.Save(filename)
}
func convertNewJsonToXmind(newJson *NewJson) *Content {
	if newJson == nil || len(*newJson) == 0 {
		return nil
	}

	// Create Content struct
	content := contentInit()

	// Process each mind map
	for _, mindMap := range *newJson {
		// Create a new sheet
		sheet := &XMLSheet{
			Title: mindMap.Title,
			Id:    mindMap.ID,
			Theme: themeid,
		}

		// Process root topic
		rootTopic := mindMap.RootTopic
		sheetTopic := Topic{
			Id:             rootTopic.ID,
			Title:          rootTopic.Title,
			Structureclass: rootTopic.StructureClass,
		}
		// Process markers
		if len(rootTopic.Markers) > 0 {
			var markerref []MarkerRef
			for i := 0; i < len(rootTopic.Markers); i++ {
				markerref = append(markerref, MarkerRef{rootTopic.Markers[i].MarkerId})

			}
			sheetTopic.MarkerRefs.MarkerRef = markerref
		}

		// Process child topics
		if len(rootTopic.Childrennew.Attached) > 0 {
			children := &Children{
				Topics: Topics{
					Type: "attached",
				},
			}

			for _, attached := range rootTopic.Childrennew.Attached {
				childTopic := Topic{
					Id:    attached.ID,
					Title: attached.Title,
				}
				children.Topics.Topic = append(children.Topics.Topic, childTopic)
			}

			sheetTopic.Children = children
		}

		sheet.Topic = sheetTopic
		content.XMLSheets = append(content.XMLSheets, sheet)
	}

	return content
}
