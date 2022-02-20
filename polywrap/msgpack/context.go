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

func (c Context) Pop() string {
	if c.IsEmpty() {
		panic("Null pointer exception: tried to pop an item from an empty Context stack")
	}

	node, a := c.nodes[len(c.nodes)-1], c.nodes[:len(c.nodes)-1]
	c.nodes = a

	nodeInfo := ""
	if node.nodeInfo != "" {
		nodeInfo = node.nodeInfo
	}
	return node.nodeItem + ": " + node.nodeType + nodeInfo
}
