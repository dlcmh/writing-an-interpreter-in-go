package lexer

import "monkey/token"

type Lexer struct {
	/*
	  The ones that might cause some confusion right now are position and readPosition.
	  Both will be used to access characters in input by using them as an index,
	  e.g.: l.input[l.readPosition]. The reason for these two "pointers" pointing into
	  our input string is the fact that we will need to be able to "peek" further into
	  the input and look after the current character to see what comes up next.
	  readPosition always points to the "next" character in the input.
	  position points to the character in the input that corresponds to the ch byte.
	*/
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination - only supports ASCII,
	// to support Unicode & UTF-8, needs to be changed to rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	/*
		The purpose of readChar is to give us the next character and advance our position
		in the input string. The first thing it does is to check whether we reached
		the end of input and can't read anymore characters. If that's the case it sets
		l.ch to 0, which is the ASCII code for the "NUL" character and signifies either
		"we haven't read anything yet" or "end of file" for us. But if we haven't reached
		the end of input yet it sets l.ch to the next character by accessing
		l.input[l.readPosition].

		After that l.position is updated to the just used l.readPosition and l.readPosition
		is incremented by one. That way, l.readPosition always points to the next position
		where we're going to read from next and l.position always points to the position
		where we last read. This will come in handy soon enough.
	*/
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
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
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar() // advance our pointers into the input so when we call NextToken()
	// again the l.ch field is already updated
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
