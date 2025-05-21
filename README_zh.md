# goxmind

`goxmind` 是一个用 Go 语言编写的库，用于创建、操作和保存 XMind 格式的思维导图文件。方便构建思维导图的结构，并将其保存为标准的 XMind 文件。

## 功能特性
- **思维导图结构创建**：支持创建 XMind 文件、添加画布、根节点和子节点。
- **节点属性设置**：可以为节点添加图标、备注、链接。
- **json文件加载**：支持按照定义好的json结构生成xmind文件。

## 兼容性
- **支持的 XMind 版本**：已测试通过可支持 旧版xmind8 和 Xmind (2025) ，如果有其他版本需求请联系我。


## 安装指南
### 环境要求
- Go 语言环境（建议 1.18 及以上版本）


### 安装步骤
可以在 `go.mod` 文件中添加如下引用：
```bash
go get github.com/vivian0517/goxmind
```

## 使用方法
### 示例代码
```go
package main

import "github.com/vivian0517/goxmind"

func main() {
	//初始化
	xmind := goxmind.NewXmind()
	//添加画布名称和根节点名称
	rootNode := xmind.AddSheet("画布名称", "根节点名称")

	//添加子节点名称
	child1 := rootNode.AddTopic("Child 1") //如果要在此节点下继续添加节点需要保存返回值
	rootNode.AddTopic("Child 2")           //如果不在此节点下继续添加节点，可以忽略返回值
	rootNode.AddTopic("Child 3")
	rootNode.AddTopic("Child 4")

	//在child1节点下继续添加子节点
	child1_1 := child1.AddTopic("Child 1.1") //如果要在此节点下继续添加节点需要保存返回值
	child1_2 := child1.AddTopic("Child 1.2")
	child1.AddTopic("Child 1.3") //如果不在此节点下继续添加节点，可以忽略返回值

	//在child1.1节点下继续添加子节点
	child1_1.AddTopic("Child 1.1.1")
	child1_1.AddTopic("Child 1.1.2")
	child1_2.AddTopic("Child 1.2.1")
	child1_2.AddTopic("Child 1.2.2")

	//保存xmind,".xmind"文件后缀可填也可不填
	xmind.Save("xmind文件名")
}

```

### 运行示例
将上述代码保存为 `main.go`，然后在终端中运行：
```bash
go run main.go
```

## 贡献指南
如果你希望为 `goxmind` 项目做出贡献，请遵循以下步骤：
1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 将更改推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 许可证
本项目采用 [MIT 许可证](LICENSE)。

## 联系方式
如果你有任何问题或建议，请在 GitHub 仓库中提交 issue。 
