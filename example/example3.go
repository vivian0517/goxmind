package main

import "github.com/vivian0517/goxmind"

func example3() {
	// Initialize
	xmind := goxmind.New()
	// Add sheet title and root node title
	rootNode := xmind.AddSheet("Sheet title", "Root Node title")

	// Add child node title
	child1 := rootNode.AddNode("Child 1") // If you want to add icons, hyperlinks, etc. to this node, you need to save the return value
	// Set hyperlink for child1 node
	child1.AddHref("www.example.com")

	child2 := rootNode.AddNode("Child 2")
	// Set notes for child2 node
	child2.AddNotes("Notes")

	child3 := rootNode.AddNode("Child 3")
	// Set icon for child3 node 🔢 Priority
	child3.AddMarker(goxmind.Priority1)
	child3.AddMarker(goxmind.Priority2)

	child4 := rootNode.AddNode("Child 4")
	// Set icon for child4 node ⭐ Star
	child4.AddMarker(goxmind.StarRed)
	// Set icon for child4 node 😊 Smiley
	child4.AddMarker(goxmind.SmileySmile)
	// Set icon for child4 node ✅ Task Progress
	child4.AddMarker(goxmind.Task0_8)
	// For more icons, refer to the MarkerId constant in marker.go

	// Save xmind, the ".xmind" file suffix is optional
	xmind.Save("xmind_file_name")
}
