package main

func example2() {

	//初始化
	xmind := x.NewXmind()
	//添加画布名称和根节点名称
	rootNode := xmind.AddSheet("画布名称", "根节点名称")

	//添加子节点名称
	child1 := xmind.AddTopic(rootNode, "Child 1") //如果要在此节点下继续添加节点需要保存返回值
	xmind.AddTopic(rootNode, "Child 2")           //如果不在此节点下继续添加节点，可以忽略返回值
	xmind.AddTopic(rootNode, "Child 3")
	xmind.AddTopic(rootNode, "Child 4")

	//在child1节点下继续添加子节点
	child1_1 := xmind.AddTopic(child1, "Child 1.1")
	child1_2 := xmind.AddTopic(child1, "Child 1.2")
	xmind.AddTopic(child1, "Child 1.3")

	//在child1.1节点下继续添加子节点
	xmind.AddTopic(child1_1, "Child 1.1.1")
	xmind.AddTopic(child1_1, "Child 1.1.2")
	xmind.AddTopic(child1_2, "Child 1.2.1")
	xmind.AddTopic(child1_2, "Child 1.2.2")

	//保存xmind,".xmind"文件后缀可填也可不填
	xmind.Save("xmind文件名")
}
