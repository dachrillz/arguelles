package interpreter

import (
	"strings"
	"testing"
)

var data = `
English French Description Gender Number ;
I       je 	   (desc)	   ()	  ()     ;
you     tu	   (mul, desc) ()     (s)    ;
he      il     ()          ()     ()     ;
she     elle   ()          ()     ()     ;
we      nous   ()          ()     ()     ;
you     vous   ()          ()     (p)    ;
they    ils    () 	       (m)    ()     ;
they    elles  ()          (f)    ()     ;
`

func assertLexer(t *testing.T, index int, l *lexer, exValue string, exType typ) {
	token := l.tokens[index]
	err := false
	if token.value != exValue {
		err = true
		t.Errorf("Mismatch in value: %s, expeted %s at %d", token.value, exValue, index)
	}
	if token.typ != exType {
		err = true
		t.Errorf("Mismatch in type: %v, expeted %v %d", token.typ, exType, index)
	}

	if err {
		t.Logf("Token raw values: %s %v", l.tokens[index].value, l.tokens[index].typ)
	}
}

func TestLex(t *testing.T) {
	l := NewLexer()
	input := strings.NewReader(data)
	l.lex(input)

	assertLexer(t, 0, l, "English", word)
	assertLexer(t, 1, l, "French", word)
	assertLexer(t, 2, l, "Description", word)
	assertLexer(t, 3, l, "Gender", word)
	assertLexer(t, 4, l, "Number", word)
	assertLexer(t, 5, l, ";", newLine)

	assertLexer(t, 6, l, "I", word)
	assertLexer(t, 7, l, "je", word)
	assertLexer(t, 8, l, "(", leftParen)
	assertLexer(t, 9, l, "desc", word)
	assertLexer(t, 10, l, ")", rightParen)
	assertLexer(t, 11, l, "(", leftParen)
	assertLexer(t, 12, l, ")", rightParen)
	assertLexer(t, 13, l, "(", leftParen)
	assertLexer(t, 14, l, ")", rightParen)
	assertLexer(t, 15, l, ";", newLine)

	assertLexer(t, 16, l, "you", word)
	assertLexer(t, 17, l, "tu", word)
	assertLexer(t, 18, l, "(", leftParen)
	assertLexer(t, 19, l, "mul", word)
	assertLexer(t, 20, l, ",", comma)
	assertLexer(t, 21, l, "desc", word)
	assertLexer(t, 22, l, ")", rightParen)
	assertLexer(t, 23, l, "(", leftParen)
	assertLexer(t, 24, l, ")", rightParen)
	assertLexer(t, 25, l, "(", leftParen)
	assertLexer(t, 26, l, "s", word)
	assertLexer(t, 27, l, ")", rightParen)
	assertLexer(t, 28, l, ";", newLine)

	//he      il     ()          ()     ()     ;
	assertLexer(t, 29, l, "he", word)
	assertLexer(t, 30, l, "il", word)
	assertLexer(t, 31, l, "(", leftParen)
	assertLexer(t, 32, l, ")", rightParen)
	assertLexer(t, 33, l, "(", leftParen)
	assertLexer(t, 34, l, ")", rightParen)
	assertLexer(t, 35, l, "(", leftParen)
	assertLexer(t, 36, l, ")", rightParen)
	assertLexer(t, 37, l, ";", newLine)

	//she     elle   ()          ()     ()     ;
	assertLexer(t, 38, l, "she", word)
	assertLexer(t, 39, l, "elle", word)
	assertLexer(t, 40, l, "(", leftParen)
	assertLexer(t, 41, l, ")", rightParen)
	assertLexer(t, 42, l, "(", leftParen)
	assertLexer(t, 43, l, ")", rightParen)
	assertLexer(t, 44, l, "(", leftParen)
	assertLexer(t, 45, l, ")", rightParen)
	assertLexer(t, 46, l, ";", newLine)

	//we      nous   ()          ()     ()     ;
	assertLexer(t, 47, l, "we", word)
	assertLexer(t, 48, l, "nous", word)
	assertLexer(t, 49, l, "(", leftParen)
	assertLexer(t, 50, l, ")", rightParen)
	assertLexer(t, 51, l, "(", leftParen)
	assertLexer(t, 52, l, ")", rightParen)
	assertLexer(t, 53, l, "(", leftParen)
	assertLexer(t, 54, l, ")", rightParen)
	assertLexer(t, 55, l, ";", newLine)

	//you     vous   ()          ()     (p)    ;
	assertLexer(t, 56, l, "you", word)
	assertLexer(t, 57, l, "vous", word)
	assertLexer(t, 58, l, "(", leftParen)
	assertLexer(t, 59, l, ")", rightParen)
	assertLexer(t, 60, l, "(", leftParen)
	assertLexer(t, 61, l, ")", rightParen)
	assertLexer(t, 62, l, "(", leftParen)
	assertLexer(t, 63, l, "p", word)
	assertLexer(t, 64, l, ")", rightParen)
	assertLexer(t, 65, l, ";", newLine)

	//they    ils    () 	       (m)    ()     ;
	assertLexer(t, 66, l, "they", word)
	assertLexer(t, 67, l, "ils", word)
	assertLexer(t, 68, l, "(", leftParen)
	assertLexer(t, 69, l, ")", rightParen)
	assertLexer(t, 70, l, "(", leftParen)
	assertLexer(t, 71, l, "m", word)
	assertLexer(t, 72, l, ")", rightParen)
	assertLexer(t, 73, l, "(", leftParen)
	assertLexer(t, 74, l, ")", rightParen)
	assertLexer(t, 75, l, ";", newLine)

	//they    elles  ()          (f)    ()     ;
	assertLexer(t, 76, l, "they", word)
	assertLexer(t, 77, l, "elles", word)
	assertLexer(t, 78, l, "(", leftParen)
	assertLexer(t, 79, l, ")", rightParen)
	assertLexer(t, 80, l, "(", leftParen)
	assertLexer(t, 81, l, "f", word)
	assertLexer(t, 82, l, ")", rightParen)
	assertLexer(t, 83, l, "(", leftParen)
	assertLexer(t, 84, l, ")", rightParen)
	assertLexer(t, 85, l, ";", newLine)

	assertLexer(t, 86, l, "eof", eof)
}
