package main

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) current() Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return Token{Type: EOF}
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) Parse() *Program {
	var statements []Node
	for p.current().Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}
	return &Program{Statements: statements}
}

func (p *Parser) parseStatement() Node {
	switch p.current().Type {
	case DESPICABLE:
		return p.parseVarDecl()
	case STAMPA:
		return p.parseFuncCall()
	case POKA:
		return p.parseIfStatement()
	case WHILEO:
		return p.parseWhileStatement()
	case WORKYO:
		return p.parseFunctionDef()
	case TORNADU:
		return p.parseReturnStmt()
	case POR:
		return p.parseForLoop()
	case SMASH:
		p.advance()
		return &BreakStmt{}
	case AVANTI:
		p.advance()
		return &ContinueStmt{}
	default:
		p.advance()
		return nil
	}
}

func (p *Parser) parseVarDecl() Node {
	p.advance()
	name := p.current().Value
	p.advance()
	p.advance()
	value := p.parseExpr()
	return &VarDecl{Name: name, Value: value}
}

func (p *Parser) parseFuncCall() Node {
	funcName := p.current().Value
	p.advance() // skip function name
	p.advance() // skip "("
	var args []Node
	for p.current().Type != RPAREN && p.current().Type != EOF {
		args = append(args, p.parseExpr())
		if p.current().Type == RPAREN {
			break
		}
		p.advance() // move to next arg
	}
	p.advance() // skip ")"
	return &FuncCall{Name: funcName, Args: args}
}

func (p *Parser) parseIfStatement() Node {
	p.advance()
	condition := p.parseExpr()
	p.advance()
	var body []Node
	for p.current().Type != RBRACE {
		stmt := p.parseStatement()
		if stmt != nil {
			body = append(body, stmt)
		}
	}
	p.advance()

	var elseBody []Node
	if p.current().Type == NADABOBA {
		p.advance()
		p.advance() // skip "{"
		for p.current().Type != RBRACE {
			stmt := p.parseStatement()
			if stmt != nil {
				elseBody = append(elseBody, stmt)
			}
		}
		p.advance()
	}

	return &IfStatement{Condition: condition, Body: body, Else: elseBody}
}

func (p *Parser) parseWhileStatement() Node {
	p.advance()
	condition := p.parseExpr()
	p.advance()
	var body []Node
	for p.current().Type != RBRACE {
		stmt := p.parseStatement()
		if stmt != nil {
			body = append(body, stmt)
		}
	}
	p.advance()
	return &WhileLoop{Condition: condition, Body: body}
}

func (p *Parser) parseList() Node {
	p.advance() // skip [
	var elements []Node
	for p.current().Type != RBRACKET && p.current().Type != EOF {
		elements = append(elements, p.parseExpr())
		if p.current().Type == COMMA {
			p.advance()
		}
	}
	p.advance() // skip ]
	return &ListLit{Elements: elements}
}

func (p *Parser) parseFunctionDef() Node {
	p.advance() // skip "workyo"
	name := p.current().Value
	p.advance() // skip name
	p.advance() // skip "("

	var params []string
	for p.current().Type != RPAREN {
		params = append(params, p.current().Value)
		p.advance()
		if p.current().Type == COMMA {
			p.advance()
		}
	}
	p.advance() // skip ")"
	p.advance() // skip "{"

	var body []Node
	for p.current().Type != RBRACE {
		stmt := p.parseStatement()
		if stmt != nil {
			body = append(body, stmt)
		}
	}
	p.advance() // skip "}"

	return &FunctionDef{Name: name, Params: params, Body: body}
}

func (p *Parser) parseReturnStmt() Node {
	p.advance() // skip "tornadu"
	value := p.parseExpr()
	return &ReturnStmt{Value: value}
}

func (p *Parser) parseForLoop() Node {
	p.advance() // skip "por"
	varName := p.current().Value
	p.advance() // skip var
	p.advance() // skip "="
	start := p.parseExpr()
	p.advance() // skip "todu"
	end := p.parseExpr()
	p.advance() // skip "{"

	var body []Node
	for p.current().Type != RBRACE {
		stmt := p.parseStatement()
		if stmt != nil {
			body = append(body, stmt)
		}
	}
	p.advance() // skip "}"

	return &ForLoop{Variable: varName, Start: start, End: end, Body: body}
}

func (p *Parser) parseExpr() Node {
	return p.parseComparison()
}

func (p *Parser) parseComparison() Node {
	left := p.parseLogicalOr()
	for p.current().Type == GT || p.current().Type == LT || p.current().Type == EQEQ {
		op := p.current().Value
		p.advance()
		right := p.parseLogicalOr()
		left = &BinaryOp{Left: left, Op: op, Right: right}
	}
	return left
}

// parser.go — update all the boolean logic
func (p *Parser) parseLogicalOr() Node {
	left := p.parseLogicalAnd()
	for p.current().Type == RO {
		op := p.current().Value
		p.advance()
		right := p.parseLogicalAnd()
		left = &BinaryOp{Left: left, Op: op, Right: right}
	}
	return left
}

func (p *Parser) parseLogicalAnd() Node {
	left := p.parseUnary()
	for p.current().Type == E {
		op := p.current().Value
		p.advance()
		right := p.parseUnary()
		left = &BinaryOp{Left: left, Op: op, Right: right}
	}
	return left
}

func (p *Parser) parseUnary() Node {
	if p.current().Type == NOD {
		p.advance()
		right := p.parseUnary()
		return &UnaryOp{Op: "nod", Right: right}
	}
	return p.parseAddition()
}

func (p *Parser) parseAddition() Node {
	left := p.parseMultiplication()
	for p.current().Type == PLUS || p.current().Type == MINUS {
		op := p.current().Value
		p.advance()
		right := p.parseMultiplication()
		left = &BinaryOp{Left: left, Op: op, Right: right}
	}
	return left
}

func (p *Parser) parseMultiplication() Node {
	left := p.parsePrimary()
	for p.current().Type == STAR || p.current().Type == SLASH {
		op := p.current().Value
		p.advance()
		right := p.parsePrimary()
		left = &BinaryOp{Left: left, Op: op, Right: right}
	}
	return left
}

func (p *Parser) parsePrimary() Node {
	tok := p.current()
	switch tok.Type {
	case NUMBER:
		p.advance()
		return &Number{Value: tok.Value}
	case IDENT:
		name := tok.Value
		p.advance()
		// Check for array indexing
		if p.current().Type == LBRACKET {
			p.advance() // skip "["
			index := p.parseExpr()
			p.advance() // skip "]"
			return &IndexAccess{List: &Identifier{Name: name}, Index: index}
		}
		// Check if it's a function call
		if p.current().Type == LPAREN {
			p.advance() // skip "("
			var args []Node
			for p.current().Type != RPAREN && p.current().Type != EOF {
				args = append(args, p.parseExpr())
				if p.current().Type == COMMA {
					p.advance()
				}
			}
			p.advance() // skip ")"
			return &FuncCall{Name: name, Args: args}
		}
		return &Identifier{Name: name}
	case SIDU:
		p.advance()
		return &BoolLit{Value: true}
	case NONO:
		p.advance()
		return &BoolLit{Value: false}
	case LBRACKET:
		return p.parseList()
	case LPAREN:
		p.advance()
		expr := p.parseExpr()
		p.advance()
		return expr
	case STRINGO:
		p.advance()
		return &StringLit{Value: tok.Value}
	default:
		return nil
	}
}
