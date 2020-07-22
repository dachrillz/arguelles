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
	VocabMap      map[string]Vocab
	headerMap     map[int]header
	languageIndex int
	stack         []Token
	vocabStack    []TargetVocab
}

func NewParser() *Parser {
	return &Parser{
		VocabMap:  make(map[string]Vocab, 0),
		headerMap: make(map[int]header, 0),
		stack:     make([]Token, 0),
	}
}

func (p *Parser) emitHeader() {
	for i, token := range p.stack {
		if i == 0 {
			p.headerMap[p.languageIndex] = *newHeader(token.value)
		} else {
			p.headerMap[p.languageIndex].headers[i] = token.value
		}
	}
	p.stack = make([]Token, 0)
}

func (p *Parser) buildHeaders(tokens []Token) (int, error) {
	var i int
	for i, token := range tokens {
		if token.typ == newLine {
			p.emitHeader()
			p.languageIndex = 0
			return i + 1, nil
		}
		if token.typ == bar {
			p.emitHeader()
			p.languageIndex++
			continue
		} else if token.typ != word {
			return i + 1, fmt.Errorf("Header was not simple word but was: '%s'", token.value)
		}
		p.stack = append(p.stack, token)
	}
	p.languageIndex = 0
	return i + 1, errors.New("Reached end of file before being able to construct headers")
}

func (p *Parser) pop() Token {
	var x Token
	x, p.stack = p.stack[len(p.stack)-1], p.stack[:len(p.stack)-1]
	return x
}

func (p *Parser) emitLine() {
	var glossary string
	for i, vocabPart := range p.vocabStack {
		if i == 0 {
			glossary = vocabPart.Target
			p.VocabMap[glossary] = Vocab{Source: glossary, TargetMap: make(map[string]TargetVocab, 0)}
		} else {
			p.VocabMap[glossary].TargetMap[vocabPart.Language] = vocabPart
		}
	}
	p.languageIndex = 0
	p.vocabStack = make([]TargetVocab, 0)
}

func (p *Parser) emitPartialVocab() {

	for i := len(p.stack); i < 4; i++ {
		p.stack = append(p.stack, Token{value: ""})
	}

	p.vocabStack = append(p.vocabStack,
		TargetVocab{
			Language: p.headerMap[p.languageIndex].language,
			Target:   p.stack[0].value,
			Desc:     p.stack[1].value,
			Gender:   p.stack[2].value,
			Number:   p.stack[3].value,
		})
	p.stack = make([]Token, 0)
}

func (p *Parser) combineParantheses(tokens []Token) (int, error) {
	var res = ""
	for i, token := range tokens {
		if token.typ == leftParen {
			continue
		} else if token.typ == rightParen {
			p.stack = append(p.stack, Token{value: res, typ: word})
			return i, nil
		} else if token.typ == eof {
			return i, errors.New("Reaced end of file before being able to combine parentheses")
		} else {
			res += token.value
		}
	}
	return -1, errors.New("Reaced end of file before being able to combine parentheses")
}

func (p *Parser) parseLines(tokens []Token) error {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].typ == bar {
			p.emitPartialVocab()
			p.languageIndex++
		} else if tokens[i].typ == leftParen {
			res, err := p.combineParantheses(tokens[i:])
			i += res
			if err != nil {
				return err
			}
		} else if tokens[i].typ == newLine {
			p.emitPartialVocab()
			p.emitLine()
		} else {
			p.stack = append(p.stack, tokens[i])
		}

	}

	println("Jere")
	for _, v := range p.VocabMap {
		fmt.Printf("%v\n", v)
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
