package internal

import (
	"fmt"
	"unicode"
)

// TokenType == May be better as an integer!
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Pos     int
}

const (
	ILLEGAL = "Illegal"
	EOF     = "eof"

	CLASS = "class"

	INDEX      = "index"
	UNIQUE     = "unique"
	SPARSE     = "sparse"
	BACKGROUND = "background"
	ASC        = "asc"
	DESC       = "desc"

	CBLOCK = "~"

	ENTITY = "entity"

	ENUM = "enum"

	AS         = "as"
	DATA       = "data"
	EQUAL      = "="
	IMPORT     = "import"
	PACKAGE    = "package"
	LBRACE     = "{"
	LPAREN     = "("
	RPAREN     = ")"
	SEMI       = ";"
	RBRACE     = "}"
	COMMA      = ","
	IDENTIFIER = "ident"
	SHOW       = "show"
)

type Lexer struct {
	input        string // The complete input
	runes        []rune
	FName        string
	lineNum      int  // The line number
	lPos         int  // Position of token on the line
	position     int  // the current character
	readPosition int  // The next position
	ch           rune // current character
}

func NewLexer(input string, fname string) *Lexer {
	l := &Lexer{input: input, FName: fname, lineNum: 1, lPos: 0, runes: []rune(input)}
	l.readChar() // Prime the first character
	return l
}

// readChar
// places current character in l.ch
// advances the pointer to the next character
func (l *Lexer) readChar() {

	if l.isEOF() {
		return
	}

	if l.readPosition >= len(l.runes) {
		l.ch = 0
	} else {
		l.ch = l.runes[l.readPosition]
		if l.ch == '\n' {
			l.lineNum += 1
			l.lPos = 0
		}
	}
	l.position = l.readPosition
	l.lPos += 1
	l.readPosition += 1
}
func (l *Lexer) isEOL() bool {
	return l.ch == '\n'
}
func (l *Lexer) isEOF() bool {
	return l.readPosition > len(l.runes)
}
func (l *Lexer) peekCharAt(idx int) rune {
	if l.readPosition >= len(l.runes)-idx {
		return 0
	} else {
		return l.runes[l.readPosition+idx]
	}
}
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.runes) {
		return 0
	} else {
		return l.runes[l.readPosition]
	}
}

func (l *Lexer) newToken(tokenType TokenType, ch rune) Token {
	var r = Token{Type: tokenType, Literal: string(ch), Line: l.lineNum, Pos: l.lPos}
	return r
}

func (l *Lexer) newTokenStr(tokenType TokenType, ch string) Token {
	var r = Token{Type: tokenType, Literal: ch, Line: l.lineNum, Pos: l.lPos}
	return r
}

// NextToken
// On input  l.ch is the current character
// On output should be at next character after token required characters.
// peekChar will be looking at the next character
func (l *Lexer) NextToken() (tk *Token) {

	defer func() {
		if err := recover(); err != nil {
			tk1 := l.newTokenStr(ILLEGAL, fmt.Sprintf("Error parsing: %s", err))
			tk = &tk1
		}
	}()

	var tok Token

	if l.isEOF() {
		tok := l.newToken(EOF, '0')
		return &tok
	}

	for {
		l.bypassWhiteSpace()
		if l.ch == '/' && l.peekChar() == '*' {
			l.bypassMultilineComment()
		} else {
			break
		}

	}

	if l.isEOF() {
		tok := l.newToken(EOF, '0')
		return &tok
	}

	switch l.ch {
	case '~':
		tok = l.newTokenStr(CBLOCK, l.readCodeBlock())
		break
	case '=':
		tok = l.newToken(EQUAL, '=')
		l.readChar()
		break
	case '(':
		tok = l.newToken(LPAREN, '(')
		l.readChar()
		break
	case ')':
		tok = l.newToken(RPAREN, ')')
		l.readChar()
		break
	case '{':
		tok = l.newToken(LBRACE, '{')
		l.readChar()
		break
	case '}':
		tok = l.newToken(RBRACE, '}')
		l.readChar()
		break
	case ',':
		tok = l.newToken(COMMA, ',')
		l.readChar()
		break
	case ';':
		tok = l.newToken(SEMI, ';')
		l.readChar()
		break
	default:
		tok = l.newTokenStr(IDENTIFIER, l.readIdentifier())
	}

	// Should always be on the next character
	if tok.Type == IDENTIFIER {
		if len(tok.Literal) == 0 {
			tok = l.newTokenStr(ILLEGAL, fmt.Sprintf("Error parsing"))
		} else {
			tok = l.convertToKeyword(tok)
		}
	}

	return &tok

}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for l.isIdentChar() {
		l.readChar()
		if l.isEOF() {
			break
		}
	}

	if l.isEOF() {
		// panic(fmt.Sprintf("Expected any of '%s' not found.", string(chars)))
		return string(l.runes[pos:l.position])
	}

	// l.readChar()
	return string(l.runes[pos:l.position])
}

// @* ... *@
func (l *Lexer) bypassMultilineComment() {
	l.readChar() // Get past the *
	for !l.isEOF() {
		if l.ch == '*' {
			if l.peekChar() == '/' {
				l.readChar()
				l.readChar()
				return
			}
		}
		l.readChar()
	}
}

func (l *Lexer) bypassWhiteSpace() {
	for !l.isEOF() {
		if unicode.IsSpace(l.ch) {
			l.readChar()
			continue
		}
		return
	}
}

func (l *Lexer) isIdentChar() bool {
	if unicode.IsDigit(l.ch) {
		return true
	}

	if unicode.IsLetter(l.ch) {
		return true
	}

	if l.ch == '.' {
		return true
	}

	if l.ch == '_' {
		return true
	}

	return false

}

func (l *Lexer) convertToKeyword(tok Token) Token {

	switch tok.Literal {
	case "entity":
		tok.Type = ENTITY
		break
	case "show":
		tok.Type = SHOW
		break
	case "index":
		tok.Type = INDEX
		break
	case "unique":
		tok.Type = UNIQUE
		break
	case "background":
		tok.Type = BACKGROUND
		break
	case "sparse":
		tok.Type = SPARSE
		break
	case "asc":
		tok.Type = ASC
		break
	case "desc":
		tok.Type = DESC
		break

	case "class":
		tok.Type = CLASS
		break
	case "enum":
		tok.Type = ENUM
		break
	case "as":
		tok.Type = AS
		break
	case "data":
		tok.Type = DATA
		break
	case "import":
		tok.Type = IMPORT
		break
	case "package":
		tok.Type = PACKAGE
		break
	case "ident":
		tok.Type = IDENTIFIER
		break
	}
	return tok
}

// readCodeBlock
//
//	~text[]anything~
func (l *Lexer) readCodeBlock() string {
	l.readChar()
	pos := l.position
	for l.ch != '~' {
		l.readChar()
		if l.isEOF() {
			panic("unexpected EOF, unterminated ~ block")
		}
	}

	str := string(l.runes[pos:l.position])
	l.readChar()
	return str
}

func (l *Lexer) PeekToken() *Token {

	cp := &Lexer{
		input:        l.input,
		runes:        l.runes,
		FName:        l.FName,
		lineNum:      l.lineNum,
		lPos:         l.lPos,
		position:     l.position,
		readPosition: l.readPosition,
		ch:           l.ch,
	}
	return cp.NextToken()
}
