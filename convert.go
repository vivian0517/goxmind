// Package goxmind provides functionality for working with XMind files.
package goxmind

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// contentInit initializes a new Content struct with default values.
func contentInit() *Content {
	return &Content{
		Modified:   "vivian",
		Xmlns:      "urn:xmind:xmap:xmlns:content:2.0",
		Xmlnsfo:    "http://www.w3.org/1999/XSL/Format",
		Xmlnssvg:   "http://www.w3.org/2000/svg",
		Xmlnsxhtml: "http://www.w3.org/1999/xhtml",
		Xmlnsxlink: "http://www.w3.org/1999/xlink",
	}
}

// MarshalContent marshals an Xmind struct into a Content struct.
func MarshalContent(xmind *Xmind) *Content {
	content := contentInit()
	for i := 0; i < len(xmind.Sheets); i++ {
		newSheet := &XMLSheet{
			Topic: convertNodeToTopic(&xmind.Sheets[i].Node),
			Title: xmind.Sheets[i].SheetTitle,
			Id:    RandStringRunes(26),
			Theme: themeid,
		}
		content.XMLSheets = append(content.XMLSheets, newSheet)
	}
	return content
}

// lettersAndDigits is a constant string containing all possible characters for random string generation.
const lettersAndDigits = "abcdefghijklmnopqrstuvwxyz0123456789"

// RandStringRunes generates a random string of the specified length.
func RandStringRunes(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Use the current time as the seed
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(lettersAndDigits[rand.Intn(len(lettersAndDigits))])
	}
	return sb.String()
}

// convertNodeToTopic converts a Node struct to a Topic struct.
func convertNodeToTopic(node *Node) Topic {
	topic := Topic{
		Title: node.NodeTitle,
		Id:    RandStringRunes(26),
	}

	// Set icons
	if len(node.Markers) > 0 {
		topic.MarkerRefs = &MarkerRefs{
			MarkerRef: make([]MarkerRef, len(node.Markers)),
		}
		for i, marker := range node.Markers {
			topic.MarkerRefs.MarkerRef[i] = MarkerRef{MarkerId: marker.Marker}
		}
	}

	// Set notes
	if node.Notes != "" {
		topic.Notes = &Note{}
		topic.Notes.Plain = node.Notes
	}

	// Set hyperlink
	if node.Href != "" {
		topic.Href = node.Href
	}

	if len(node.Children) > 0 {
		//增加extensions
		num := strconv.Itoa(len(node.Children))
		topic.Extensions = &Extensions{
			Extension: Extension{
				Provider: provider,
				Ext_content: Ext_content{
					Right_number: num,
				},
			},
		}

		children := Children{
			Topics: Topics{
				Type:  "attached",
				Topic: []Topic{},
			},
		}
		for _, child := range node.Children {
			children.Topics.Topic = append(children.Topics.Topic, convertNodeToTopic(child))
		}
		topic.Children = &children
	}

	return topic
}
