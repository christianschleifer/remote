package terminal

import (
	"github.com/ChristianSchleifer/mremoteng/pkg/controller/api"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type terminalViewer struct {
	controller api.Controller
}

func NewViewer(controller api.Controller) api.Viewer {
	return &terminalViewer{
		controller,
	}
}

func (terminalView *terminalViewer) View() {
	collection := terminalView.controller.GetCollection()
	root := tview.NewTreeNode(collection.Name).SetExpanded(true).SetSelectable(false)

	terminalView.add(root, collection.Children[0])

	treeView := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	application := tview.NewApplication()
	application.SetRoot(treeView, true)
	application.EnableMouse(true)

	if err := application.Run(); err != nil {
		panic(err)
	}
}

func (terminalView *terminalViewer) add(target *tview.TreeNode, node api.Node) {
	var newNode *tview.TreeNode
	switch n := node.(type) {
	case *api.Collection:
		newNode = tview.NewTreeNode(n.Name)
		newNode.SetColor(tcell.ColorLightGray)
		newNode.SetSelectable(true)
		newNode.SetExpanded(false)

		newNode.SetSelectedFunc(func() {
			newNode.SetExpanded(!newNode.IsExpanded())
		})

		for _, child := range n.Children {
			terminalView.add(newNode, child)
		}

	case *api.Connection:
		newNode = tview.NewTreeNode(n.Name)
		newNode.SetColor(tcell.ColorDarkGreen)
		newNode.SetSelectable(true)
		newNode.SetExpanded(false)

		newNode.SetSelectedFunc(func() {
			terminalView.controller.ConnectionSelectedHandler(n.Id)
		})
	}

	target.AddChild(newNode)
}
