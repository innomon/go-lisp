package lisp

import (
	"fmt"
	"testing"
)

func equalSlices(a, b Tokens) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.val != b[i].val || v.typ != b[i].typ {
			return false
		}
	}
	return true
}

func TestNewTokens(t *testing.T) {
	var tests = []struct {
		in  string
		out Tokens
	}{
		{"(define a 42)", Tokens{{openToken, "("}, {symbolToken, "define"}, {symbolToken, "a"}, {numberToken, "42"}, {closeToken, ")"}}},
		{"\t(quote\n\t\t(a b c))  ", Tokens{{openToken, "("}, {symbolToken, "quote"}, {openToken, "("}, {symbolToken, "a"}, {symbolToken, "b"}, {symbolToken, "c"}, {closeToken, ")"}, {closeToken, ")"}}},
		{"hello ; dude\n\tworld", Tokens{{symbolToken, "hello"}, {symbolToken, "world"}}},
		{"test \"a string\"", Tokens{{symbolToken, "test"}, {stringToken, "\"a string\""}}},
		{"\"only string\"", Tokens{{stringToken, "\"only string\""}}},
		{"\"string\\nwith\\\"escape\\tcharacters\"", Tokens{{stringToken, "\"string\\nwith\\\"escape\\tcharacters\""}}},
		{"\"hej\\\"hello\"", Tokens{{stringToken, "\"hej\\\"hello\""}}},
		{"'hej", Tokens{{symbolToken, "'"}, {symbolToken, "hej"}}},
	}

	for _, test := range tests {
		x := NewTokens(test.in)
		if !equalSlices(x, test.out) {
			t.Errorf("NewTokens \"%v\" gives \"%v\", expected \"%v\"", test.in, x, test.out)
		}
	}
}

func TestParse(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"42", "(42)"},
		{"(+ (+ 1 2) 3)", "((+ (+ 1 2) 3))"},
	}
	for _, test := range tests {
		parsed, _ := NewTokens(test.in).Parse()
		result := fmt.Sprintf("%v", parsed.Inspect())
		if result != test.out {
			t.Errorf("Parse \"%v\" gives \"%v\", expected \"%v\"", test.in, result, test.out)
		}
	}
}

func TestParseFailures(t *testing.T) {
	var tests = []string{
		"(42",
	}
	for _, in := range tests {
		if x, err := NewTokens(in).Parse(); err == nil {
			t.Errorf("Parse('%v') = '%v', want error", in, x)
		}
	}
}
