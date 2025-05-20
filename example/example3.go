package main

import "github.com/vivian0517/goxmind"

func example3() {
	//初始化
	xmind := goxmind.New()
	//添加画布名称和根节点名称
	rootNode := xmind.AddSheet("画布名称", "根节点名称")

	//添加子节点名称
	child1 := rootNode.AddNode("Child 1") //如果要在此节点增加图标超链接等需要保存返回值
	//给child1节点设置超链接
	child1.AddHref("www.example.com")

	child2 := rootNode.AddNode("Child 2")
	//给child2节点设置备注
	child2.AddNotes("备注")

	child3 := rootNode.AddNode("Child 3")
	//给child3节点设置图标 🔢 优先级
	child3.AddMaker(goxmind.Priority1)
	child3.AddMaker(goxmind.Priority2)

	child4 := rootNode.AddNode("Child 4")
	//给child4节点设置图标 ⭐ 星星
	child4.AddMaker(goxmind.StarRed)
	//给child4节点设置图标 😊 表情
	child4.AddMaker(goxmind.SmileySmile)
	//给child4节点设置图标 ✅ 任务进度
	child4.AddMaker(goxmind.Task0_8)
	//更多图标参考marker.go中MarkerId常量

	//保存xmind,".xmind"文件后缀可填也可不填
	xmind.Save("xmind文件名")
}
