package token

// string not as performant as byte or int, but has convenient methods for
// debugging, eg we could just print a string
type TokenType string

// the different token types in the Monkey language
const (
  ILLEGAL = "ILLEGAL" // a token or character we don't know about
  EOF = "EOF" // end of file - tells our parser later on that it can stop

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y, ...
  INT = "INT" // 123456

  // Operators
  ASSIGN = "="
  PLUS = "+"

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)

// Token data structure
type Token struct {
  Type TokenType // To distinguish between eg "integer", "right bracket", etc
  Literal string // Holds the literal value of the token
}
