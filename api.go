package xml

const (
	// Header A generic XML header
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// Pos the position in the original input text
type Pos int

// Node a node of parsed XML document
type Node interface {
	Parent() Node
	SetParent(p Node) Node
	Tag() string
	Attrs() map[string]string
	Children() []Node
	Position() Pos
}

// Document a parsed XML document
type Document interface {
	// The root of the document
	Root() Node
	SetRoot(r Node)
}

// Parse converts a xml document into a Document
func Parse(s string) (Document, error) {
	return nil, nil
}

// Serialize converts a Document into xml text
func Serialize(doc Document) (string, error) {
	return "", nil
}
