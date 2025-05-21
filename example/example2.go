package main

import "github.com/vivian0517/goxmind"

func example2() {

	// Initialize
	xmind := goxmind.New()
	// Add sheet title and root node title
	rootNode := xmind.AddSheet("Sheet title", "Root Node title")

	// Add child node title
	child1 := xmind.AddNode(rootNode, "Child 1") // If you want to add more nodes under this node, you need to save the return value
	xmind.AddNode(rootNode, "Child 2")           // If you don't want to add more nodes under this node, you can ignore the return value
	xmind.AddNode(rootNode, "Child 3")
	xmind.AddNode(rootNode, "Child 4")

	// Continue adding child nodes under the child1 node
	child1_1 := xmind.AddNode(child1, "Child 1.1")
	child1_2 := xmind.AddNode(child1, "Child 1.2")
	xmind.AddNode(child1, "Child 1.3")

	// Continue adding child nodes under the child1.1 node
	xmind.AddNode(child1_1, "Child 1.1.1")
	xmind.AddNode(child1_1, "Child 1.1.2")
	xmind.AddNode(child1_2, "Child 1.2.1")
	xmind.AddNode(child1_2, "Child 1.2.2")

	// Save xmind, the ".xmind" file suffix is optional
	xmind.Save("xmind_file_name")
}
