package interpreter

import (
	"fmt"
	"testing"
)

func assertHeaders(t *testing.T, p *Parser) {
	fmt.Println(p.headerMap)
	assertHeader(t, 0, p, "English", []string{})
	assertHeader(t, 1, p, "French", []string{"Description", "Gender", "Number"})
	assertHeader(t, 2, p, "Spanish", []string{"Description", "Gender", "Number"})
}

func assertHeader(t *testing.T, index int, p *Parser, expectedLanguage string, expected []string) {
	if p.headerMap[index].language != expectedLanguage {
		t.Errorf("Header at %d was '%s', but expected '%s'", index, p.headerMap[index].language, expectedLanguage)
	}

	for i, header := range p.headerMap[index].headers {
		if header != expected[i-1] {
			t.Errorf("Header for %s at %d was '%s', but expected '%s'", p.headerMap[index].language, i, header, expected[i-1])
		}
	}
}

func TestParser(t *testing.T) {
	//English | French Description Gender Number  | Spanish Description Gender Number ;
	//I       | je 	   (desc)	   ()	  ()      | yo 	    () 			() 	   ()     ;
	//you     | tu	   (mul, desc) ()     (s)     | tu      ()          ()     ()	  ;
	tokens := []Token{

		// Headers
		{value: "English", typ: word},
		{value: "|", typ: bar},

		{value: "French", typ: word},
		{value: "Description", typ: word},
		{value: "Gender", typ: word},
		{value: "Number", typ: word},
		{value: "|", typ: bar},

		{value: "Spanish", typ: word},
		{value: "Description", typ: word},
		{value: "Gender", typ: word},
		{value: "Number", typ: word},
		{value: ";", typ: newLine},

		// I
		{value: "I", typ: word},
		{value: "|", typ: bar},

		{value: "je", typ: word},
		{value: "(", typ: leftParen},
		{value: "desc", typ: word},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: "|", typ: bar},

		{value: "yo", typ: word},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: ";", typ: newLine},

		// you
		{value: "you", typ: word},
		{value: "|", typ: bar},

		{value: "tu", typ: word},
		{value: "(", typ: leftParen},
		{value: "mul", typ: word},
		{value: ",", typ: comma},
		{value: "desc", typ: word},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: "s", typ: word},
		{value: ")", typ: rightParen},
		{value: "|", typ: bar},

		{value: "tu", typ: word},
		{value: "(", typ: leftParen},
		{value: ",", typ: comma},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: ")", typ: rightParen},
		{value: "(", typ: leftParen},
		{value: "s", typ: word},
		{value: ")", typ: rightParen},
		{value: ";", typ: newLine},

		{value: "eof", typ: eof},
	}

	p := NewParser()

	//fmt.Printf("%v\n", p.headerMap)

	if err := p.Parse(tokens); err != nil {
		t.Error(err)
	} else {
		assertHeaders(t, p)
	}
}
