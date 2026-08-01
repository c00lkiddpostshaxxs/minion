package main

type Lexer struct {
	input string
	pos   int
	ch    byte
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input, pos: 0}
	if len(input) > 0 {
		lex.ch = input[0]
	}
	return lex
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos < len(l.input) {
		l.ch = l.input[l.pos]
	} else {
		l.ch = 0
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' {
		l.advance()
	}
	// Skip comments
	if l.ch == '#' {
		for l.ch != '\n' && l.ch != 0 {
			l.advance()
		}
		l.skipWhitespace() // recurse, don't return
	}
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for l.ch >= '0' && l.ch <= '9' {
		l.advance()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readIdent() string {
	start := l.pos
	for isAlpha(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.advance()
	}
	return l.input[start:l.pos]
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readString() string {
	l.advance() // skip opening quote
	start := l.pos
	for l.ch != '"' && l.ch != 0 {
		l.advance()
	}
	result := l.input[start:l.pos]
	l.advance() // skip closing quote
	return result
}

func lookupIdent(ident string) TokenType {
	switch ident {
	case "despicable":
		return DESPICABLE
	case "stampa":
		return STAMPA
	case "poka":
		return POKA
	case "whileo":
		return WHILEO
	case "sidu":
		return SIDU
	case "nono":
		return NONO
	case "workyo":
		return WORKYO
	case "tornadu":
		return TORNADU
	case "nadaboba":
		return NADABOBA
	case "e":
		return E
	case "ro":
		return RO
	case "nod":
		return NOD
	case "por":
		return POR
	case "todu":
		return TODU
	case "smash":
		return SMASH
	case "avanti":
		return AVANTI
	default:
		return IDENT
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token

	switch l.ch {
	case '(':
		tok = Token{Type: LPAREN, Value: "("}
	case ')':
		tok = Token{Type: RPAREN, Value: ")"}
	case '{':
		tok = Token{Type: LBRACE, Value: "{"}
	case '}':
		tok = Token{Type: RBRACE, Value: "}"}
	case '=':
		if l.pos+1 < len(l.input) && l.input[l.pos+1] == '=' {
			tok = Token{Type: EQEQ, Value: "=="}
			l.advance() // advance extra for the second =
		} else {
			tok = Token{Type: EQUALS, Value: "="}
		}
	case '+':
		tok = Token{Type: PLUS, Value: "+"}
	case '-':
		tok = Token{Type: MINUS, Value: "-"}
	case '*':
		tok = Token{Type: STAR, Value: "*"}
	case '/':
		tok = Token{Type: SLASH, Value: "/"}
	case '>':
		tok = Token{Type: GT, Value: ">"}
	case '<':
		tok = Token{Type: LT, Value: "<"}
	case '[':
		tok = Token{Type: LBRACKET, Value: "["}
	case ']':
		tok = Token{Type: RBRACKET, Value: "]"}
	case ',':
		tok = Token{Type: COMMA, Value: ","}
	case '"':
		str := l.readString()
		return Token{Type: STRINGO, Value: str}
	case 0:
		tok = Token{Type: EOF, Value: ""}
	default:
		if isDigit(l.ch) {
			num := l.readNumber()
			return Token{Type: NUMBER, Value: num}
		} else if isAlpha(l.ch) {
			ident := l.readIdent()
			return Token{Type: lookupIdent(ident), Value: ident}
		}
	}

	l.advance()
	return tok
}
