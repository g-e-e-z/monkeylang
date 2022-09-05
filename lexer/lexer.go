package lexer

import "monkeylang/token"

type Lexer struct {
    input           string
    position        int     // current poision in input (points to current char)
    readPosition    int     // current reading position in input (after current)
    ch              byte    // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

// Note: Lexer only supports ASCII instead of full Unicode
// To add support, we would change l.ch from byte to rune and change
// our method of reading the next characters
// cince they could be multiply bytes wide
// TODO: Add full Unicode and Emoji support
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0    // ASCII code for "NUL" -> EOF
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}
