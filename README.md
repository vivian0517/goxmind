# goxmind
Read this document in other languages:
- [简体中文 (Chinese)](README_zh.md)

`goxmind` is a library written in Go, designed for creating, manipulating, and saving mind map files in the XMind format. It simplifies the process of constructing mind map structures and saving them as standard XMind files.

## Features
- **Mind Map Structure Creation**: Supports creating XMind files, adding sheets, root nodes, and child nodes.
- **Node Attribute Settings**: Allows adding icons, notes, and links to nodes.
- **JSON File Loading**: Enables generating XMind files based on predefined JSON structures.

## Compatibility
- **Supported XMind Versions**: Tested and compatible with XMind 8 and XMind (2025). If you need support for other versions, please contact me.

## Installation Guide
### Environment Requirements
- Go programming environment (version 1.18 or higher is recommended).

### Installation Steps
You can add the following reference to your `go.mod` file:
```bash
go get github.com/vivian0517/goxmind
```

## Usage
### Example Code
```go
package main

import "github.com/vivian0517/goxmind"

func main() {
	// Initialize
	xmind := goxmind.NewXmind()
	// Add a sheet name and a root node name
	rootNode := xmind.AddSheet("Sheet Name", "Root Node Name")

	// Add child node names
	child1 := rootNode.AddTopic("Child 1") // Save the return value if you want to add more nodes under this node
	rootNode.AddTopic("Child 2")           // You can ignore the return value if you don't need to add more nodes under this node
	rootNode.AddTopic("Child 3")
	rootNode.AddTopic("Child 4")

	// Add more child nodes under child1
	child1_1 := child1.AddTopic("Child 1.1") // Save the return value if you want to add more nodes under this node
	child1_2 := child1.AddTopic("Child 1.2")
	child1.AddTopic("Child 1.3") // You can ignore the return value if you don't need to add more nodes under this node

	// Add more child nodes under child1.1
	child1_1.AddTopic("Child 1.1.1")
	child1_1.AddTopic("Child 1.1.2")
	child1_2.AddTopic("Child 1.2.1")
	child1_2.AddTopic("Child 1.2.2")

	// Save the XMind file. The ".xmind" file extension is optional.
	xmind.Save("xmind_file_name")
}
```

### Running the Example
Save the above code as `main.go`, then run the following command in the terminal:
```bash
go run main.go
```
### Result
After running the example, you will find a file named `xmind_file_name.xmind` in the same directory. This file is the generated XMind file.
![generated XMind](./example/1.png)

## Features
```go
package main

import "github.com/vivian0517/goxmind"

func main() {
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
	child3.AddMaker(goxmind.Priority1)
	child3.AddMaker(goxmind.Priority2)

	child4 := rootNode.AddNode("Child 4")
	// Set icon for child4 node ⭐ Star
	child4.AddMaker(goxmind.StarRed)
	// Set icon for child4 node 😊 Smiley
	child4.AddMaker(goxmind.SmileySmile)
	// Set icon for child4 node ✅ Task Progress
	child4.AddMaker(goxmind.Task0_8)
	// For more icons, refer to the MarkerId constant in marker.go

	// Save xmind, the ".xmind" file suffix is optional
	xmind.Save("xmind_file_name")
}
```
![generated XMind](./example/2.png)

## Contribution Guide
If you wish to contribute to the `goxmind` project, please follow these steps:
1. Fork this repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push the changes to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## License
This project is licensed under the [MIT License](LICENSE).

## Contact
If you have any questions or suggestions, please submit an issue in the GitHub repository.
---


        