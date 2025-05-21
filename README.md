# goxmind

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
Read this document in other languages:
- [简体中文 (Chinese)](README_zh.md)
```

        