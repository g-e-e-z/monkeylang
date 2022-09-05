package token

// Declaring as a string allows us to use many different values
// TokenTypes, which allows us to ditinguis between different types
// of tokens.
// Using string also has the advantage of being easy to debug without
// a lot of boilerplate and helper functions, we can just print a string
// Of course using a string might not lead to the samem performance as
// using an int or a byte would.

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifers + literlas
    IDENT   = "IDENT"   // add, foobar, x, y, ...
    INT     = "INT"     // 2135235

    // Operators
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"

    LT      = "<"
    GT      = ">"

    EQ      = "=="
    NOT_EQ  = "!="

    // Delimiters
    COMMA       = ","
    SEMICOLON   = ";"

    LPAREN  = "("
    RPAREN  = ")"
    LBRACE  = "{"
    RBRACE  = "}"

    // Keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
    IF          = "IF"
    ELSE        = "ELSE"
    RETURN      = "RETURN"
)

var keywords = map[string]TokenType{
    "fn":       FUNCTION,
    "let":      LET,
    "true":     TRUE,
    "false":    FALSE,
    "if":       IF,
    "else":     ELSE,
    "return":   RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

