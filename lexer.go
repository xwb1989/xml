package xml

import (
	"unicode/utf8"
)

type itemType int
type stateFn func(*lexer) stateFn

// Lexer that tokenize a xml document
type lexer struct {
	input string
	items chan item
	state stateFn
	pos   Pos
	width Pos
	start Pos
}

const (
	itemStartTag itemType = iota
	itemEndTag
	itemEmptyTag
	itemContent
	itemAttrVal
	itemAttrID
	itemEOF
	itemErr
)

const (
	leftAngle    = '<'
	rightAngle   = '>'
	forwardSlash = '/'
)
const eof = -1
const invalidEncoding = -2

type item struct {
	typ itemType
	val string
	pos Pos
}

func lex(s string) *lexer {
	l := &lexer{
		items: make(chan item),
	}
	go l.run()
	return l
}

// Main lexing logic defined here
func (l *lexer) run() {
	for l.state = lexText; l.state != nil; l.state = l.state(l) {
	}
	close(l.items)

}

func lexText(l *lexer) stateFn {
	l.skipWhiteSpaces()
	if l.peek() == leftAngle {
		return lexOpenTag
	}
	return errorRecovery
}

func lexOpenTag(l *lexer) stateFn {
	l.skipWhiteSpaces()
	for {
		r := l.next()
		switch r {
		case rightAngle:
			// End of start tag
		case forwardSlash:
			// Empty tag
		default:

		}
	}
	return nil
}

func errorRecovery(l *lexer) stateFn {
	return nil
}

func (l *lexer) skipWhiteSpaces() {
	for r := l.next(); isWhiteSpace(r); r = l.next() {
	}
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	if r == utf8.RuneError {
		if w == 0 {
			return eof
		}
		return invalidEncoding
	}
	l.width = Pos(w)
	l.pos += l.width
	return rune(r)
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
