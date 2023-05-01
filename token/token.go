package token

// 定义标识符和关键字常量
const (
	ILLEGAL = "ILLEGAL" // 未知的词法单元或字符
	EOF     = "EOF"     // 表示文件结尾

	// 标识符 + 字面量
	IDENT = "IDENT" // 用户当前自定义的标识符 add ,foobar, x, y
	INT   = "INT"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NotEq    = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType // token 类型
	Literal string    // token类型对应的字符
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}
	return IDENT
}
