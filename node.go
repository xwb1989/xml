package xml

type node struct {
	parent   Node
	tag      string
	attrs    map[string]string
	children []Node
	pos      Pos
}

// NewNode creates an empty xml node
func NewNode() Node {
	return &node{parent: nil, children: []Node{}, attrs: map[string]string{}}
}

func (n *node) Attrs() map[string]string {
	return n.attrs
}

func (n *node) Tag() string {
	return n.tag
}

func (n *node) SetTag(s string) Node {
	n.tag = s
	return n
}

func (n *node) Position() Pos {
	return n.pos
}

func (n *node) Children() []Node {
	return n.children
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) SetParent(p Node) Node {
	n.parent = p
	return n
}

func (n *node) SetChildren(c []Node) Node {
	n.children = c
	return n
}
