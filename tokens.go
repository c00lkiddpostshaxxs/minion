package main

type TokenType string

const (
	DESPICABLE TokenType = "DESPICABLE"
	STAMPA     TokenType = "STAMPA"
	WORKYO     TokenType = "WORKYO"
	TORNADU    TokenType = "TORNADU"
	POKA       TokenType = "POKA"
	WHILEO     TokenType = "WHILEO"
	SIDU       TokenType = "SIDU"
	NONO       TokenType = "NONO"
	NADABOBA   TokenType = "NADABOBA"
	NUMBER     TokenType = "NUMBER"
	IDENT      TokenType = "IDENT"
	LPAREN     TokenType = "("
	RPAREN     TokenType = ")"
	LBRACE     TokenType = "{"
	RBRACE     TokenType = "}"
	EQUALS     TokenType = "="
	PLUS       TokenType = "+"
	MINUS      TokenType = "-"
	STAR       TokenType = "*"
	SLASH      TokenType = "/"
	GT         TokenType = ">"
	LT         TokenType = "<"
	EQEQ       TokenType = "=="
	EOF        TokenType = "EOF"
	LBRACKET   TokenType = "["
	RBRACKET   TokenType = "]"
	COMMA      TokenType = ","
	STRINGO    TokenType = "STRINGO"
	E          TokenType = "E"
	RO         TokenType = "RO"
	NOD        TokenType = "NOD"
	POR        TokenType = "POR"
	TODU       TokenType = "TODU"
	SMASH      TokenType = "SMASH"
	AVANTI     TokenType = "AVANTI"
)

type Token struct {
	Type  TokenType
	Value string
}
