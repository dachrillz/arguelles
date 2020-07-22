package interpreter

import (
	"io"
	"strings"
	"unicode"
)

type typ int
type spaceState int

const (
	word typ = iota
	leftParen
	rightParen
	bar
	newLine
	comma
	eof
)

const (
	start spaceState = iota
	notSpace
	space
)

type lexer struct {
	tokens []Token
	state  spaceState
}

type Token struct {
	typ   typ
	value string
}

func NewLexer() *lexer {
	return &lexer{}
}

func isWhiteSpace(c rune) bool {
	return unicode.IsSpace(c)
}

func (l *lexer) isSingleCharacter(c rune, buf *[]string) bool {
	switch c {
	case '(':
		l.tokens = append(l.tokens, Token{value: "(", typ: leftParen})
		return true
	case '|':
		if len(*buf) > 0 {
			l.addToken(strings.Join(*buf, ""), word)
			*buf = make([]string, 0)
		}
		l.tokens = append(l.tokens, Token{value: "|", typ: bar})
		return true
	case ')':
		if len(*buf) > 0 {
			l.addToken(strings.Join(*buf, ""), word)
			*buf = make([]string, 0)
		}
		l.tokens = append(l.tokens, Token{value: ")", typ: rightParen})
		return true
	case ';':
		l.tokens = append(l.tokens, Token{value: ";", typ: newLine})
		return true
	case ',':
		if len(*buf) > 0 {
			l.addToken(strings.Join(*buf, ""), word)
			*buf = make([]string, 0)
		}
		l.tokens = append(l.tokens, Token{value: ",", typ: comma})
		return true
	default:
		return false
	}
}

func (l *lexer) addToken(value string, typ typ) {
	l.tokens = append(l.tokens, Token{value: value, typ: typ})
}

func (l *lexer) lexStream(stream string) {
	buf := make([]string, 0)
	for _, c := range stream {
		if l.isSingleCharacter(c, &buf) {
			continue
		}
		if isWhiteSpace(c) {
			if l.state == space {
				continue
			} else if l.state == notSpace {
				l.state = space
				if len(buf) > 0 {
					l.addToken(strings.Join(buf, ""), word)
					buf = make([]string, 0)
				}
			}
		} else {
			l.state = notSpace
			buf = append(buf, string(c))
		}
	}
	l.addToken("eof", eof)
}

func (l *lexer) lex(r io.Reader) ([]Token, error) {
	//@TODO: Convert to concurrent emission of tokens :)

	buf := make([]byte, 32*1024) // define your buffer size here.
	for {

		n, err := r.Read(buf)

		if n > 0 {
			l.lexStream(string(buf[:n])) // your read buffer.
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return make([]Token, 0), nil

}
