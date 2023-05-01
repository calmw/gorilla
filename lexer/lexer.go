package lexer

import (
	"fmt"
	"gorilla/token"
)

type Lexer struct {
	input        string
	position     int  // 所输入字符串中的当前位置（指向当前字符串）
	readPosition int  // 所输入字符串中当前读取位置（指向当前字符串之后的一个字符）
	ch           byte // 当前正在查看的字符，如果要完全支持UTF8，这里类型要使用rune,还需要修改读取下一个字符的方式
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

// 读取input中的下一个字符，并前移其在input中的位置
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 窥视输入中的下一个字符。与readChar相比，不会前移l.position和l.readPosition
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NotEq,
				Literal: literal,
			}
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case 0: // 读到最后一个字符，再往后读值为0
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if l.isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			fmt.Println(1)
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()

	return tok
}

// 可以控制变量中允许的字符
func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 这里做了简化处理，只支持整数，没有对浮点数、16进制数、8进制数等的支持
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		fmt.Println("skip: ", string(l.ch))
		l.readChar()
	}
}

// 可以将readNumber和readIdentifier合并为一个函数来处理，但为了便于理解，分开处理
