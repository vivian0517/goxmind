package main

import "github.com/vivian0517/goxmind"

func example1() {

	// Initialize
	xmind := goxmind.New()
	// Add sheet title and root node title
	rootNode := xmind.AddSheet("Sheet title", "Root Node title")

	// Add child node title
	child1 := rootNode.AddNode("Child 1") // If you want to add more nodes under this node, you need to save the return value
	rootNode.AddNode("Child 2")           // If you don't want to add more nodes under this node, you can ignore the return value
	rootNode.AddNode("Child 3")
	rootNode.AddNode("Child 4")

	// Continue adding child nodes under the child1 node
	child1_1 := child1.AddNode("Child 1.1") // If you want to add more nodes under this node, you need to save the return value
	child1_2 := child1.AddNode("Child 1.2")
	child1.AddNode("Child 1.3") // If you don't want to add more nodes under this node, you can ignore the return value

	// Continue adding child nodes under the child1.1 node
	child1_1.AddNode("Child 1.1.1")
	child1_1.AddNode("Child 1.1.2")
	child1_2.AddNode("Child 1.2.1")
	child1_2.AddNode("Child 1.2.2")

	// Save xmind, the ".xmind" file suffix is optional
	xmind.Save("xmind_file_name")

}
