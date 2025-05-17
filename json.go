package goxmind

import (
	"errors"
	"strconv"
)

type Xmind struct {
	FileName string   `json:"filename"`
	Canvas   []*Sheet `json:"sheet"`
}

type Sheet struct {
	SheetTitle string `json:"sheetTitle"`
	Node       Node   `json:"node"`
}

type Node struct {
	NodeTitle string   `json:"nodeTitle"`
	Children  []*Node  `json:"children"`
	Makers    []Makers `json:"makers"`
	Notes     string   `json:"notes"`
	Href      string   `json:"href"`
}

type Makers struct {
	Maker string `json:"maker"`
}

func NewXmind() *Xmind {

	return &Xmind{}
}

// 修改文件名
func (x *Xmind) SetFileName(filename string) {
	x.FileName = filename
}

// 添加xmind的一个新画布和根节点，传参画布名称，根节点名称，返回根节点指针
func (x *Xmind) AddSheet(sheetTitle string, rootTitle string) *Node {
	canva := &Sheet{
		SheetTitle: sheetTitle,
		Node: Node{
			NodeTitle: rootTitle,
		},
	}
	x.Canvas = append(x.Canvas, canva)
	return &x.Canvas[len(x.Canvas)-1].Node
}

// 添加xmind的一个新子节点，传参父节点指针，要添加的子节点名称，返回新添加的节点指针
func (x *Xmind) AddTopic(node *Node, str string) *Node {
	child := &Node{
		NodeTitle: str,
	}
	node.Children = append(node.Children, child)
	return node.Children[len(node.Children)-1]
}

// 通过node 添加一个新的子节点，传参要添加的子节点名称，返回新添加的节点指针
func (n *Node) AddTopic(str string) *Node {
	child := &Node{
		NodeTitle: str,
	}
	n.Children = append(n.Children, child)
	return n.Children[len(n.Children)-1]
}

// 在node!之前!添加一个新的节点，传参要添加的节点名称，返回新添加的节点指针
func (x *Xmind) AddTopicBefore(node *Node, str string) (*Node, error) {
	if node == nil {
		return nil, errors.New("node is nil")
	}

	for _, sheet := range x.Canvas {
		parent := FindParent(&sheet.Node, node)
		if parent != nil {
			// 在 parent.Children 中找到 node 并替换为新 parent 节点
			newParent := &Node{
				NodeTitle: str,
				Children:  []*Node{node},
			}
			for i, child := range parent.Children {
				if child == node {
					parent.Children[i] = newParent
					return newParent, nil
				}
			}
		}
	}

	return nil, errors.New("can't find parent node")
}

func FindParent(root, target *Node) *Node {
	if root == nil || target == nil {
		return nil
	}
	for _, child := range root.Children {
		if child == target {
			return root
		}
		if found := FindParent(child, target); found != nil {
			return found
		}
	}
	return nil
}

func (x *Xmind) DeleteNode(node *Node) error {
	if node == nil {
		return errors.New("node is nil")
	}

	for _, sheet := range x.Canvas {
		parent := FindParent(&sheet.Node, node)
		if parent != nil {
			// 在 parent.Children 中找到 node 并删除
			newChildren := []*Node{}
			for _, child := range parent.Children {
				if child != node {
					newChildren = append(newChildren, child)
				}
			}
			parent.Children = newChildren
			return nil
		} else {
			return errors.New("not found parent node")
		}
	}

	return nil
}

// 更新xmind的一个节点名字，传节点指针，要更新的Title
func (x *Xmind) UpdateTopic(node *Node, str string) {
	node.NodeTitle = str
}

// 更新xmind的一个节点名字，传节点指针，要更新的Title
func (n *Node) UpdateTopic(str string) {
	n.NodeTitle = str
}

type ProrityLevel int

const (
	ProrityOne   ProrityLevel = 1
	ProrityTwo   ProrityLevel = 2
	ProrityThree ProrityLevel = 3
	ProrityFour  ProrityLevel = 4
	ProrityFive  ProrityLevel = 5
	ProritySix   ProrityLevel = 6
	ProritySeven ProrityLevel = 7
	ProrityEight ProrityLevel = 8
	ProrityNine  ProrityLevel = 9
)

// 设置图标
func (x *Xmind) AddMaker(node *Node, maker ProrityLevel) {
	prostr := "priority-" + strconv.Itoa(int(maker))
	makerval := Makers{
		Maker: prostr,
	}
	node.Makers = append(node.Makers, makerval)
}

// 设置图标
func (n *Node) AddMaker(maker ProrityLevel) {
	prostr := "priority-" + strconv.Itoa(int(maker))
	makerval := Makers{
		Maker: prostr,
	}
	n.Makers = append(n.Makers, makerval)
}

// 设置备注
func (x *Xmind) AddNotes(node *Node, notes string) {
	node.Notes = notes
}

// 设置备注
func (n *Node) AddNotes(notes string) {
	n.Notes = notes
}

// 设置超链接
func (x *Xmind) AddHref(node *Node, href string) {
	node.Href = href
}

// 设置超链接
func (n *Node) AddHref(href string) {
	n.Href = href
}
