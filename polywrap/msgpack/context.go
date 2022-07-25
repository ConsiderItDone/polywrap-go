package msgpack

type Node struct {
	nodeItem string
	nodeType string
	nodeInfo string
}

type Context struct {
	description string
	nodes       []Node
}

func NewContext(description string) *Context {
	return &Context{description: description}
}

func (c *Context) IsEmpty() bool {
	return len(c.nodes) == 0
}

func (c *Context) Length() int32 {
	return int32(len(c.nodes))
}

func (c *Context) Push(nodeItem, nodeType, nodeInfo string) {
	c.nodes = append(c.nodes, Node{
		nodeItem: nodeItem,
		nodeType: nodeType,
		nodeInfo: nodeInfo,
	})
}

func (c *Context) Pop() string {
	if c.IsEmpty() {
		panic("Null pointer exception: tried to pop an item from an empty Context stack")
	}

	node, a := c.nodes[len(c.nodes)-1], c.nodes[:len(c.nodes)-1]
	c.nodes = a

	nodeInfo := ""
	if node.nodeInfo != "" {
		nodeInfo = " >> " + node.nodeInfo
	}
	return node.nodeItem + ": " + node.nodeType + nodeInfo
}

func (c *Context) toString() string {
	return c.printWithTabs(0, 2)
}

func (c *Context) PrintWithContext(message string) string {
	return message + "\n" + c.printWithTabs(1, 2)
}

func (c *Context) printWithTabs(tabs, size int32) string {
	result := lpad("", " ", size*tabs)
	result += "Context: " + c.description
	tabs++

	if c.IsEmpty() {
		result += rpad("\n", " ", size*tabs+1)
		result += "context stack is empty"

		return result
	}

	for i := len(c.nodes) - 1; i >= 0; i-- {
		node := c.nodes[i]
		nodeInfo := ""
		if node.nodeInfo != "" {
			nodeInfo = " >> " + node.nodeInfo
		}

		result += rpad("\n", " ", size*tabs+1)
		tabs++
		result += "at " + node.nodeItem + ": " + node.nodeType + nodeInfo
	}

	return result
}

func lpad(s string, pad string, plength int32) string {
	for i := int32(len(s)); i < plength; i++ {
		s = pad + s
	}

	return s
}

func rpad(s string, pad string, plength int32) string {
	for i := int32(len(s)); i < plength; i++ {
		s = s + pad
	}

	return s
}
