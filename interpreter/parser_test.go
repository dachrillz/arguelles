package interpreter

import (
	"testing"
)

func assertHeaders(t *testing.T, p *Parser) {
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

func assertContent(t *testing.T, p *Parser) {
	vocab := p.VocabMap["I"]

	if vocab.Source != "I" {
		t.Error("Wrong source")
	}

	// French
	r := vocab.TargetMap["French"].Desc
	if r != "desc" {
		t.Errorf("Wrong '%s' for French should be: desc", r)
	}
	r = vocab.TargetMap["French"].Gender
	if r != "" {
		t.Errorf("Wrong '%s' for French should be: ''", r)
	}
	r = vocab.TargetMap["French"].Language
	if r != "French" {
		t.Errorf("Wrong '%s' for French should be: 'French'", r)
	}
	r = vocab.TargetMap["French"].Number
	if r != "" {
		t.Errorf("Wrong '%s' for French should be: ''", r)
	}
	r = vocab.TargetMap["French"].Target
	if r != "je" {
		t.Errorf("Wrong '%s' for French should be: 'je'", r)
	}

	// Spanish
	r = vocab.TargetMap["Spanish"].Desc
	if r != "" {
		t.Errorf("Wrong '%s' for Spanish should be: ''", r)
	}
	r = vocab.TargetMap["Spanish"].Gender
	if r != "" {
		t.Errorf("Wrong '%s' for Spanish should be: ''", r)
	}
	r = vocab.TargetMap["Spanish"].Language
	if r != "Spanish" {
		t.Errorf("Wrong '%s' for Spanish should be: 'Spanish'", r)
	}
	r = vocab.TargetMap["Spanish"].Number
	if r != "" {
		t.Errorf("Wrong '%s' for Spanish should be: ''", r)
	}
	r = vocab.TargetMap["Spanish"].Target
	if r != "yo" {
		t.Errorf("Wrong '%s' for Spanish should be: 'yo'", r)
	}

	vocab = p.VocabMap["you"]

	if vocab.Source != "you" {
		t.Error("Wrong source")
	}

	// French
	r = vocab.TargetMap["French"].Desc
	if r != "mul,desc" {
		t.Errorf("Wrong '%s' for French should be: desc", r)
	}
	r = vocab.TargetMap["French"].Gender
	if r != "" {
		t.Errorf("Wrong '%s' for French should be: ''", r)
	}
	r = vocab.TargetMap["French"].Language
	if r != "French" {
		t.Errorf("Wrong '%s' for French should be: 'French'", r)
	}
	r = vocab.TargetMap["French"].Number
	if r != "s" {
		t.Errorf("Wrong '%s' for French should be: 's'", r)
	}
	r = vocab.TargetMap["French"].Target
	if r != "tu" {
		t.Errorf("Wrong '%s' for French should be: 'tu'", r)
	}

	// Spanish
	r = vocab.TargetMap["Spanish"].Desc
	if r != "" {
		t.Errorf("Wrong '%s' for Spanish should be: ''", r)
	}
	r = vocab.TargetMap["Spanish"].Gender
	if r != "" {
		t.Errorf("Wrong '%s' for Spanish should be: 's'", r)
	}
	r = vocab.TargetMap["Spanish"].Language
	if r != "Spanish" {
		t.Errorf("Wrong '%s' for Spanish should be: 'Spanish'", r)
	}
	r = vocab.TargetMap["Spanish"].Number
	if r != "s" {
		t.Errorf("Wrong '%s' for Spanish should be: 's'", r)
	}
	r = vocab.TargetMap["Spanish"].Target
	if r != "tu" {
		t.Errorf("Wrong '%s' for Spanish should be: 'tu'", r)
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

	if err := p.Parse(tokens); err != nil {
		t.Error(err)
	} else {
		assertHeaders(t, p)
		assertContent(t, p)
	}
}
