package xmind

import (
	"math/rand"
	"strings"
	"time"
)

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

func MarshalContent(xmind *Xmind) *Content {
	content := contentInit()
	for i := 0; i < len(xmind.Canvas); i++ {
		newSheet := &XMLSheet{
			Topic: convertNodeToTopic(&xmind.Canvas[i].Node),
			Title: xmind.Canvas[i].SheetTitle,
			Id:    RandStringRunes(26),
		}
		content.XMLSheets = append(content.XMLSheets, newSheet)
	}
	return content
}

const lettersAndDigits = "abcdefghijklmnopqrstuvwxyz0123456789"

// RandStringRunes 生成指定长度的随机字符串
func RandStringRunes(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // 使用当前时间作为种子
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(lettersAndDigits[rand.Intn(len(lettersAndDigits))])
	}
	return sb.String()
}

func convertNodeToTopic(node *Node) Topic {
	topic := Topic{
		Title: node.NodeTitle,
		Id:    RandStringRunes(26),
	}

	// 设置标记
	if len(node.Makers) > 0 {
		topic.MakerRefs = &MakerRefs{
			MakerRef: make([]MakerRef, len(node.Makers)),
		}
		for i, maker := range node.Makers {
			topic.MakerRefs.MakerRef[i] = MakerRef{MakerId: maker.Maker}
		}
	}

	//设置备注
	if node.Notes != "" {
		topic.Notes = &Note{}

		topic.Notes.Plain = node.Notes
	}

	//设置超链接
	if node.Href != "" {
		topic.Href = node.Href
	}

	if len(node.Children) > 0 {
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
