package interpreter

import (
	"errors"
	"fmt"
)

type headerState int

const (
	inHeader headerState = iota
	inContent
)

type header struct {
	language string
	headers  map[int]string
}

func newHeader(language string) *header {
	return &header{language: language, headers: make(map[int]string, 0)}
}

type Parser struct {
	VocabMap    map[string]Vocab
	headerMap   map[int]header
	stack       []Token
	headerState headerState
}

func NewParser() *Parser {
	return &Parser{
		VocabMap:  make(map[string]Vocab, 0),
		headerMap: make(map[int]header, 0),
		stack:     make([]Token, 0),
	}
}

func (p *Parser) emitHeader(index int) {
	for i, token := range p.stack {
		if i == 0 {
			p.headerMap[index] = *newHeader(token.value)
		} else {
			p.headerMap[index].headers[i] = token.value
		}
	}
	p.stack = make([]Token, 0)
}

func (p *Parser) buildHeaders(tokens []Token) (int, error) {
	index := 0
	for _, token := range tokens {
		if token.typ == newLine {
			p.emitHeader(index)
			index++
			p.headerState = inContent
			return index, nil
		}
		if token.typ == bar {
			p.emitHeader(index)
			index++
			continue
		} else if token.typ != word {
			return index, fmt.Errorf("Header was not simple word but was: '%s'", token.value)
		}
		p.stack = append(p.stack, token)
	}
	return index, errors.New("Reached end of file before being able to construct headers")
}

func (p *Parser) parseLines(tokens []Token) error {
	var index int
	for _, token := range tokens {
		if token.typ == rightParen || token.typ == word {
			index++
		}
	}

	return nil
}

func (p *Parser) Parse(tokens []Token) error {
	index, err := p.buildHeaders(tokens)
	if err != nil {
		return err
	}

	err = p.parseLines(tokens[index:])
	if err != nil {
		return err
	}

	return nil
}
