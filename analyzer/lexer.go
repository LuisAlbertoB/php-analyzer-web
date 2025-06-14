package analyzer

import "strings"

type TokenType string

const (
	Keyword     TokenType = "keyword"
	Identifier  TokenType = "identifier"
	Operator    TokenType = "operator"
	Literal     TokenType = "literal"
	Punctuation TokenType = "punctuation"
	Comment     TokenType = "comment"
)

type Token struct {
	Type  TokenType `json:"type"`
	Value string    `json:"value"`
	Line  int       `json:"line"`
	Pos   int       `json:"pos"`
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func isOperator(c byte) bool {
	operators := []byte{'+', '-', '*', '/', '=', '<', '>', '!', '&', '|', '%', '^'}
	for _, op := range operators {
		if c == op {
			return true
		}
	}
	return false
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isIdentifierChar(c byte) bool {
	return isLetter(c) || isDigit(c)
}

func isKeyword(word string) bool {
	phpKeywords := []string{"echo", "if", "else", "while", "for", "foreach", "function", "return", "class", "public", "private", "protected"}
	for _, kw := range phpKeywords {
		if word == kw {
			return true
		}
	}
	return false
}

func LexicalAnalysis(code string) []Token {
	var tokens []Token
	lines := splitLines(code)
	for lineNum, line := range lines {
		pos := 0
		for pos < len(line) {
			char := line[pos]

			switch {
			case isWhitespace(char):
				pos++
			case char == '/' && pos+1 < len(line) && line[pos+1] == '/':
				comment := line[pos:]
				tokens = append(tokens, Token{Type: Comment, Value: comment, Line: lineNum + 1, Pos: pos + 1})
				pos = len(line)
			case char == '$':
				start := pos
				pos++
				for pos < len(line) && isIdentifierChar(line[pos]) {
					pos++
				}
				tokens = append(tokens, Token{Type: Identifier, Value: line[start:pos], Line: lineNum + 1, Pos: start + 1})
			case isOperator(char):
				tokens = append(tokens, Token{Type: Operator, Value: string(char), Line: lineNum + 1, Pos: pos + 1})
				pos++
			case char == ';' || char == ',' || char == '(' || char == ')':
				tokens = append(tokens, Token{Type: Punctuation, Value: string(char), Line: lineNum + 1, Pos: pos + 1})
				pos++
			default:
				if isDigit(char) {
					start := pos
					for pos < len(line) && isDigit(line[pos]) {
						pos++
					}
					tokens = append(tokens, Token{Type: Literal, Value: line[start:pos], Line: lineNum + 1, Pos: start + 1})
				} else if isLetter(char) {
					start := pos
					for pos < len(line) && isIdentifierChar(line[pos]) {
						pos++
					}
					word := line[start:pos]
					if isKeyword(word) {
						tokens = append(tokens, Token{Type: Keyword, Value: word, Line: lineNum + 1, Pos: start + 1})
					} else {
						tokens = append(tokens, Token{Type: Identifier, Value: word, Line: lineNum + 1, Pos: start + 1})
					}
				} else {
					pos++
				}
			}
		}
	}

	return tokens
}
