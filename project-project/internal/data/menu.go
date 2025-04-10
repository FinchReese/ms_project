package data

import "github.com/jinzhu/copier"

const (
	rootParantId = 0
)

type ProjectMenu struct {
	Id         int64
	Pid        int64
	Title      string
	Icon       string
	Url        string
	FilePath   string
	Params     string
	Node       string
	Sort       int
	Status     int
	CreateBy   int64
	IsInner    int
	Values     string
	ShowSlider int
}

func (*ProjectMenu) TableName() string {
	return "ms_project_menu"
}

type ProjectMenuNode struct {
	ProjectMenu
	Children []*ProjectMenuNode
}

func ConvertMenuListToTreeList(menuList []*ProjectMenu) []*ProjectMenuNode {
	// 先把所有节点创建出来
	nodeList := []*ProjectMenuNode{}
	copier.Copy(&nodeList, menuList)
	// 找出根节点，可能有多个
	rootNodeList := []*ProjectMenuNode{}

	for _, node := range nodeList {
		if node.Pid == rootParantId {
			rootNodeList = append(rootNodeList, node)
		}
	}
	for _, node := range rootNodeList {
		buildTreeByRootNode(node, nodeList)
	}
	return rootNodeList
}

func buildTreeByRootNode(root *ProjectMenuNode, nodeList []*ProjectMenuNode) {
	getChildrenOfNode(root, nodeList)
	for _, node := range root.Children {
		buildTreeByRootNode(node, nodeList)
	}
}

func getChildrenOfNode(target *ProjectMenuNode, nodeList []*ProjectMenuNode) {
	for _, node := range nodeList {
		if node.Pid == target.Id {
			target.Children = append(target.Children, node)
		}
	}
}
