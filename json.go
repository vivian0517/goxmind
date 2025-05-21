package goxmind

import (
	"errors"
)

type Xmind struct {
	FileName string   `json:"filename"`
	Sheets   []*Sheet `json:"sheet"`
}

type Sheet struct {
	SheetTitle string `json:"sheetTitle"`
	Node       Node   `json:"node,omitempty"`
}

type Node struct {
	NodeTitle string   `json:"nodeTitle"`
	Children  []*Node  `json:"children,omitempty"`
	Makers    []Makers `json:"makers,omitempty"`
	Notes     string   `json:"notes,omitempty"`
	Href      string   `json:"href,omitempty"`
}

type Makers struct {
	Maker string `json:"maker,omitempty"`
}

func New() *Xmind {

	return &Xmind{}
}

// SetFileName modifies the filename.
func (x *Xmind) SetFileName(filename string) {
	x.FileName = filename
}

// AddSheet adds a new sheet and root node to the xmind,
// takes the sheet name and root node name as parameters,
// and returns a pointer to the root node.
func (x *Xmind) AddSheet(sheetTitle string, rootTitle string) *Node {
	canva := &Sheet{
		SheetTitle: sheetTitle,
		Node: Node{
			NodeTitle: rootTitle,
		},
	}
	x.Sheets = append(x.Sheets, canva)
	return &x.Sheets[len(x.Sheets)-1].Node
}

// AddNode adds a new child node to the xmind,
// takes a pointer to the parent node and the name of the child node to be added as parameters,
// and returns a pointer to the newly added node.
func (x *Xmind) AddNode(node *Node, str string) *Node {
	child := &Node{
		NodeTitle: str,
	}
	node.Children = append(node.Children, child)
	return node.Children[len(node.Children)-1]
}

// AddNode adds a new child node to the xmind,
func (n *Node) AddNode(str string) *Node {
	child := &Node{
		NodeTitle: str,
	}
	n.Children = append(n.Children, child)
	return n.Children[len(n.Children)-1]
}

// In node! before! adding a new node,
func (x *Xmind) AddNodeBefore(node *Node, str string) (*Node, error) {
	if node == nil {
		return nil, errors.New("node is nil")
	}

	for _, sheet := range x.Sheets {
		parent := FindParent(&sheet.Node, node)
		if parent != nil {
			// In parent.Children find node and replace it with new parent node
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

	for _, sheet := range x.Sheets {
		parent := FindParent(&sheet.Node, node)
		if parent != nil {
			// In parent.Children find node and delete
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

// Update xmind one node name, pass node pointer, to update the Title
func (x *Xmind) UpdateNode(node *Node, str string) {
	node.NodeTitle = str
}

// Update xmind one node name, pass node pointer, to update the Title
func (n *Node) UpdateNode(str string) {
	n.NodeTitle = str
}

// Set icon
func (x *Xmind) AddMarker(node *Node, marker MarkerId) {
	markerVal := Makers{
		Maker: string(marker),
	}
	node.Makers = append(node.Makers, markerVal)
}

// Set icon
func (n *Node) AddMaker(marker MarkerId) {
	markerVal := Makers{
		Maker: string(marker),
	}
	n.Makers = append(n.Makers, markerVal)
}

// Set note
func (x *Xmind) AddNotes(node *Node, notes string) {
	node.Notes = notes
}

// Set note
func (n *Node) AddNotes(notes string) {
	n.Notes = notes
}

// Set href
func (x *Xmind) AddHref(node *Node, href string) {
	node.Href = href
}

// Set href
func (n *Node) AddHref(href string) {
	n.Href = href
}
